package systemUsage

import(
	"github.com/kaimu/speedtest/providers/netflix"
	"github.com/showwin/speedtest-go/speedtest"

)

func GetBandwithSpeed() []interface{} {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	netflixServer,_ := netflix.Fetch()
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

	return []interface{}{netflixServer} //s.Latency, s.DLSpeed, s.ULSpeed,
	}

	return nil
}