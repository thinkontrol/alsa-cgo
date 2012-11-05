package main

import (
    alsa "github.com/Narsil/alsa-go"
    "fmt"
    "os"
    "flag"
)

func main(){
    rate := flag.Int("rate", 8000, "Sample rate (in Hz) of device")
    channels := flag.Int("channels", 1, "Number of channels")
    help := flag.Bool("help", false, "help")
    var streamType alsa.StreamType
    if os.Args[0] == "./aplay"{
        streamType = alsa.StreamTypePlayback
    }else{
        streamType = alsa.StreamTypeCapture
    }


    flag.Parse()

    if *help{
        flag.Usage()
        return
    }


    handle := alsa.New()
	err := handle.Open("default", streamType, alsa.ModeBlock)
	if err != nil {
		fmt.Printf("Open failed. %s", err)
	}

	handle.SampleFormat = alsa.SampleFormatS16LE
	handle.SampleRate = *rate
	handle.Channels = *channels
	err = handle.ApplyHwParams()
	if err != nil {
		fmt.Printf("SetHwParams failed. %s", err)
	}
    if err != nil{
        fmt.Println(err)
    }
    buflen := 32
    buf := make([]byte, buflen)
    for {
        os.Stdin.Read(buf)
        n, err := handle.Write(buf)
        if err != nil{
            fmt.Println(err)
        }
        if n != buflen{
            fmt.Println("Did not read whole buffer")
        }
        fmt.Printf("%s", buf)

    }
}
