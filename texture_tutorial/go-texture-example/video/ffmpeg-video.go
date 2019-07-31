package example_video

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"

	"github.com/3d0c/gmf"
)

type ffmpegVideo struct {
	swsctx         *gmf.SwsCtx
	ist            *gmf.Stream
	cc             *gmf.CodecCtx
	inputCtx       *gmf.FmtCtx
	srcVideoStream *gmf.Stream
	Frames         chan *gmf.Packet
	player         *playerStatus
}

type playerStatus struct{ flag int32 }

const (
	allImagesProcessed = 2
	playing            = 1
	noImagesAvailable  = 0
)

func (f *ffmpegVideo) Init(videoSource string, bufferSize int) (err error) {
	f.player = new(playerStatus)

	var srcFileName string
	// hard-coded path, run the app with `hover run` in the project root
	flag.StringVar(&srcFileName, "src", videoSource, "source video")
	flag.Parse()

	f.inputCtx, err = gmf.NewInputCtx(srcFileName)
	f.Frames = make(chan *gmf.Packet, bufferSize)
	if err != nil {
		return err
	}

	f.srcVideoStream, err = f.inputCtx.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	if err != nil {
		return errors.New("No video stream found in " + srcFileName + "\n")
	}

	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		return err
	}

	f.cc = gmf.NewCodecCtx(codec)

	f.cc.
		SetTimeBase(gmf.AVR{Num: 1, Den: 1}).
		SetPixFmt(gmf.AV_PIX_FMT_RGBA).
		SetWidth(f.srcVideoStream.CodecCtx().Width()).
		SetHeight(f.srcVideoStream.CodecCtx().Height())

	if codec.IsExperimental() {
		f.cc.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := f.cc.Open(nil); err != nil {
		return err
	}

	f.ist, err = f.inputCtx.GetStream(f.srcVideoStream.Index())
	if err != nil {
		return err
	}

	// convert source pix_fmt into AV_PIX_FMT_RGBA
	// which is set up by codec context above
	icc := f.srcVideoStream.CodecCtx()
	if f.swsctx, err = gmf.NewSwsCtx(icc.Width(), icc.Height(), icc.PixFmt(), f.cc.Width(), f.cc.Height(), f.cc.PixFmt(), gmf.SWS_BICUBIC); err != nil {
		return err
	}

	return nil
}

func (f *ffmpegVideo) Free() {
	f.swsctx.Free()
	f.ist.Free()
	f.inputCtx.Free()
	gmf.Release(f.cc)
	f.cc.Free()
}

func (f *ffmpegVideo) Stream(onFirstFrame func()) {
	drain := -1
	hasConsumer := false
	var wg sync.WaitGroup

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("texture_tutorial: recover: ", r)
			for len(f.Frames) > 0 { // clean the frame channel
				<-f.Frames
			}
			f.Cancel()
		}
	}()

	for {
		if drain >= 0 {
			break
		}

		pkt, err := f.inputCtx.GetNextPacket()
		if err != nil && err != io.EOF {
			if pkt != nil {
				pkt.Free()
			}
			log.Printf("error getting next packet - %s", err)
			break
		} else if err != nil && pkt == nil {
			drain = 0
		}

		if pkt != nil && pkt.StreamIndex() != f.srcVideoStream.Index() {
			continue
		}

		frames, err := f.ist.CodecCtx().Decode(pkt)
		if err != nil {
			log.Printf("Fatal error during decoding - %s\n", err)
			break
		}

		// Decode() method doesn't treat EAGAIN and EOF as errors
		// it returns empty frames slice instead. Countinue until
		// input EOF or frames received.
		if len(frames) == 0 && drain < 0 {
			continue
		}

		if frames, err = gmf.DefaultRescaler(f.swsctx, frames); err != nil {
			panic(err)
		}

		packets, err := f.cc.Encode(frames, drain)
		if err != nil {
			log.Fatalf("Error encoding - %s\n", err)
		}
		if len(packets) == 0 {
			break
		}

		for _, p := range packets {
			f.Frames <- p
			f.Play()
			if !hasConsumer {
				wg.Add(1)
				go func() {
					onFirstFrame()
					wg.Done()
				}()
				hasConsumer = true
			}
		}

		for i := range frames {
			frames[i].Free()
		}

		if pkt != nil {
			pkt.Free()
			pkt = nil
		}
	}
	f.EndOfVideo()
	for i := 0; i < f.inputCtx.StreamsCnt(); i++ {
		st, _ := f.inputCtx.GetStream(i)
		defer st.CodecCtx().Free()
		defer st.Free()
	}

	wg.Wait()
}

func (f *ffmpegVideo) Bounds() (int, int) {
	return f.cc.Width(), f.cc.Height()
}

func (f *ffmpegVideo) GetFrameRate() float64 {
	a := f.srcVideoStream.GetRFrameRate().AVR()
	return float64(a.Den) / float64(a.Num)
}

func (f *ffmpegVideo) Duration() float64 {
	return f.inputCtx.Duration()
}

func (f *ffmpegVideo) EndOfVideo() {
	f.Set(allImagesProcessed)
}

func (f *ffmpegVideo) Play() {
	f.Set(playing)
}

func (f *ffmpegVideo) Cancel() {
	f.Set(noImagesAvailable)
}

func (f *ffmpegVideo) Set(value int32) {
	atomic.StoreInt32(&(f.player.flag), value)
}

func (f *ffmpegVideo) Closed() bool {
	flag := atomic.LoadInt32(&(f.player.flag))
	if flag == allImagesProcessed && len(f.Frames) <= 1 {
		defer f.Set(noImagesAvailable)
	}
	return flag != playing && len(f.Frames) == 0
}

func (f *ffmpegVideo) HasFrameAvailable() bool {
	flag := atomic.LoadInt32(&(f.player.flag))

	if flag == allImagesProcessed || flag == playing {
		return true
	}
	return false
}
