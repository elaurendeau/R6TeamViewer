package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"strings"
	"testing"
	"github.com/elaurendeau/R6TeamViewer/core/domain"
)

type MockedRepositoryHttpHandler struct {
	mock.Mock
}

func (mockedHttpHandler *MockedRepositoryHttpHandler) Get(url string) (HttpContent, error) {
	args := mockedHttpHandler.Called(url)

	value := reflect.ValueOf(args.Get(0))
	httpContent := value.Interface().(HttpContent)

	return httpContent, args.Error(1)
}

func TestSeasonValidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	expectedSeasons := new(domain.Seasons)
	jsonSeason := "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"
	json.NewDecoder(strings.NewReader(jsonSeason)).Decode(expectedSeasons)

	seasonRepository.HttpHandler.Get(url)

	actualSeasons, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Nil(t, err)

	assert.Equal(t, actualSeasons, expectedSeasons)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestSeasonInvalidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, errors.New("Mocked error"))

	seasonRepository.HttpHandler.Get(url)

	_, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestSeasonValidButWithInvalidHttpRequestFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/seasons?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	mockedHttpContent := HttpContent{StatusCode: 400, Status: "400 BAD REQUEST", Content: "{\"seasons\":{\"4\":{\"ncsa\":{\"wins\":11,\"losses\":10,\"abandons\":1,\"season\":4,\"region\":\"ncsa\",\"ranking\":{\"rating\":2575.32307477,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.7532307477,\"stdev\":6,\"rank\":13}},\"emea\":{\"wins\":20,\"losses\":17,\"abandons\":1,\"season\":4,\"region\":\"emea\",\"ranking\":{\"rating\":2592.06724988,\"next_rating\":2600,\"prev_rating\":2500,\"mean\":25.9206724988,\"stdev\":6,\"rank\":13}}},\"5\":{\"ncsa\":{\"wins\":20,\"losses\":16,\"abandons\":0,\"season\":5,\"region\":\"ncsa\",\"ranking\":{\"rating\":2909.35896403,\"next_rating\":3100,\"prev_rating\":2900,\"mean\":29.0935896403,\"stdev\":6,\"rank\":15}},\"emea\":{\"wins\":17,\"losses\":13,\"abandons\":1,\"season\":5,\"region\":\"emea\",\"ranking\":{\"rating\":2707.75585756,\"next_rating\":2900,\"prev_rating\":2700,\"mean\":27.0775585756,\"stdev\":6,\"rank\":14}}}}}"}

	seasonRepository := new(SeasonRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	seasonRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	seasonRepository.HttpHandler.Get(url)

	_, err := seasonRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestPlayerValidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonOperator := "{\"player\":{\"username\":\"minthok\",\"platform\":\"uplay\",\"ubisoft_id\":\"fb7dde18-2052-4eb2-b732-a65c241be262\",\"indexed_at\":\"2017-01-10T04:30:28.246Z\",\"updated_at\":\"2017-02-26T03:21:46.197Z\",\"stats\":{\"ranked\":{\"has_played\":true,\"wins\":87,\"losses\":66,\"wlr\":1.318,\"kills\":826,\"deaths\":549,\"kd\":1.505,\"playtime\":179759},\"casual\":{\"has_played\":true,\"wins\":352,\"losses\":260,\"wlr\":1.354,\"kills\":2361,\"deaths\":1534,\"kd\":1.539,\"playtime\":506173},\"overall\":{\"revives\":86,\"suicides\":13,\"reinforcements_deployed\":2562,\"barricades_built\":668,\"steps_moved\":723633,\"bullets_fired\":101936,\"bullets_hit\":25741,\"headshots\":1692,\"melee_kills\":61,\"penetration_kills\":273,\"assists\":688},\"progression\":{\"level\":106,\"xp\":13844}}}}"
	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: jsonOperator}

	playerRepository := new(PlayerRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	playerRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	expectedPlayer := new(domain.Player)
	json.NewDecoder(strings.NewReader(jsonOperator)).Decode(expectedPlayer)

	playerRepository.HttpHandler.Get(url)

	actualSeasons, err := playerRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Nil(t, err)

	assert.Equal(t, actualSeasons, expectedPlayer)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestPlayerInvalidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonPlayer := "{\"player\":{\"username\":\"minthok\",\"platform\":\"uplay\",\"ubisoft_id\":\"fb7dde18-2052-4eb2-b732-a65c241be262\",\"indexed_at\":\"2017-01-10T04:30:28.246Z\",\"updated_at\":\"2017-02-26T03:21:46.197Z\",\"stats\":{\"ranked\":{\"has_played\":true,\"wins\":87,\"losses\":66,\"wlr\":1.318,\"kills\":826,\"deaths\":549,\"kd\":1.505,\"playtime\":179759},\"casual\":{\"has_played\":true,\"wins\":352,\"losses\":260,\"wlr\":1.354,\"kills\":2361,\"deaths\":1534,\"kd\":1.539,\"playtime\":506173},\"overall\":{\"revives\":86,\"suicides\":13,\"reinforcements_deployed\":2562,\"barricades_built\":668,\"steps_moved\":723633,\"bullets_fired\":101936,\"bullets_hit\":25741,\"headshots\":1692,\"melee_kills\":61,\"penetration_kills\":273,\"assists\":688},\"progression\":{\"level\":106,\"xp\":13844}}}}"
	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: jsonPlayer}

	playerRepository := new(PlayerRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	playerRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, errors.New("Mocked error"))

	playerRepository.HttpHandler.Get(url)

	_, err := playerRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestPlayerValidButWithInvalidHttpRequestFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonPlayer := "{\"player\":{\"username\":\"minthok\",\"platform\":\"uplay\",\"ubisoft_id\":\"fb7dde18-2052-4eb2-b732-a65c241be262\",\"indexed_at\":\"2017-01-10T04:30:28.246Z\",\"updated_at\":\"2017-02-26T03:21:46.197Z\",\"stats\":{\"ranked\":{\"has_played\":true,\"wins\":87,\"losses\":66,\"wlr\":1.318,\"kills\":826,\"deaths\":549,\"kd\":1.505,\"playtime\":179759},\"casual\":{\"has_played\":true,\"wins\":352,\"losses\":260,\"wlr\":1.354,\"kills\":2361,\"deaths\":1534,\"kd\":1.539,\"playtime\":506173},\"overall\":{\"revives\":86,\"suicides\":13,\"reinforcements_deployed\":2562,\"barricades_built\":668,\"steps_moved\":723633,\"bullets_fired\":101936,\"bullets_hit\":25741,\"headshots\":1692,\"melee_kills\":61,\"penetration_kills\":273,\"assists\":688},\"progression\":{\"level\":106,\"xp\":13844}}}}"
	mockedHttpContent := HttpContent{StatusCode: 400, Status: "400 BAD REQUEST", Content: jsonPlayer}

	playerRepository := new(PlayerRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	playerRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	playerRepository.HttpHandler.Get(url)

	_, err := playerRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestOperatorValidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonOperator := "{\"operator_records\":[{\"stats\":{\"played\":19,\"wins\":10,\"losses\":9,\"kills\":21,\"deaths\":11,\"playtime\":3968,\"specials\":{\"operatorpvp_doc_selfrevive\":\"0\",\"operatorpvp_doc_hostagerevive\":\"0\",\"operatorpvp_doc_teammaterevive\":\"1\"}},\"operator\":{\"name\":\"Doc\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/i06csah0j/doc.png\",\"badge\":\"https://s19.postimg.org/e5cm0ktk3/doc.png\",\"bust\":\"https://s19.postimg.org/c6dx8rbtv/doc.png\"}}},{\"stats\":{\"played\":19,\"wins\":9,\"losses\":10,\"kills\":22,\"deaths\":13,\"playtime\":4166,\"specials\":{\"operatorpvp_smoke_poisongaskill\":\"1\"}},\"operator\":{\"name\":\"Smoke\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/u4h0vj4ir/smoke.png\",\"badge\":\"https://s19.postimg.org/89rv4nt0j/smoke.png\",\"bust\":\"https://s19.postimg.org/bpspmbo2r/smoke.png\"}}},{\"stats\":{\"played\":149,\"wins\":70,\"losses\":79,\"kills\":115,\"deaths\":101,\"playtime\":30610,\"specials\":{\"operatorpvp_rook_armorboxdeployed\":\"150\",\"operatorpvp_rook_armortakenourself\":\"149\",\"operatorpvp_rook_armortakenteammate\":\"507\"}},\"operator\":{\"name\":\"Rook\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/of5doynqb/rook.png\",\"badge\":\"https://s19.postimg.org/c3lbdtacj/rook.png\",\"bust\":\"https://s19.postimg.org/i27uw5r4z/rook.png\"}}},{\"stats\":{\"played\":94,\"wins\":56,\"losses\":38,\"kills\":62,\"deaths\":51,\"playtime\":20415,\"specials\":{\"operatorpvp_echo_enemy_sonicburst_affected\":\"41\"}},\"operator\":{\"name\":\"Echo\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/vwstsrycj/echo.png\",\"bust\":\"https://s19.postimg.org/c0085o2oz/echo.png\"}}},{\"stats\":{\"played\":31,\"wins\":17,\"losses\":14,\"kills\":24,\"deaths\":22,\"playtime\":6563,\"specials\":{\"operatorpvp_caveira_interrogations\":\"0\"}},\"operator\":{\"name\":\"Capitão\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/l4qnzxtib/capitao.png\",\"bust\":\"https://s19.postimg.org/g6pnyqroj/capitao.png\"}}},{\"stats\":{\"played\":24,\"wins\":14,\"losses\":10,\"kills\":22,\"deaths\":16,\"playtime\":4603,\"specials\":{\"operatorpvp_capitao_lethaldartkills\":\"0\"}},\"operator\":{\"name\":\"Caveira\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/egu2dca03/caveira.png\",\"bust\":\"https://s19.postimg.org/km3hq9eoz/caveira.png\"}}},{\"stats\":{\"played\":104,\"wins\":58,\"losses\":46,\"kills\":104,\"deaths\":70,\"playtime\":21448,\"specials\":{\"operatorpvp_buck_kill\":\"3\"}},\"operator\":{\"name\":\"Buck\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/nprn9ptib/buck.png\",\"badge\":\"https://s19.postimg.org/4sgma7f6r/buck.png\",\"bust\":\"https://s19.postimg.org/5p47gqg83/buck.png\"}}},{\"stats\":{\"played\":666,\"wins\":382,\"losses\":284,\"kills\":704,\"deaths\":416,\"playtime\":141805,\"specials\":{\"operatorpvp_bandit_batterykill\":\"1\"}},\"operator\":{\"name\":\"Bandit\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/hnk0j852b/bandit.png\",\"badge\":\"https://s19.postimg.org/ygjv94wir/bandit.png\",\"bust\":\"https://s19.postimg.org/adb0kfqn7/bandit.png\"}}},{\"stats\":{\"played\":163,\"wins\":90,\"losses\":73,\"kills\":164,\"deaths\":106,\"playtime\":35049,\"specials\":{\"operatorpvp_jager_gadgetdestroybycatcher\":\"180\"}},\"operator\":{\"name\":\"Jäger\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/5x631uc9v/jager.png\",\"badge\":\"https://s19.postimg.org/56mjnjoo3/jager.png\",\"bust\":\"https://s19.postimg.org/nzbqcz89v/jager.png\"}}},{\"stats\":{\"played\":18,\"wins\":9,\"losses\":9,\"kills\":12,\"deaths\":12,\"playtime\":3870,\"specials\":{\"operatorpvp_iq_gadgetspotbyef\":\"178\"}},\"operator\":{\"name\":\"IQ\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/zb1thfezn/iq.png\",\"badge\":\"https://s19.postimg.org/w36ivv7hf/iq.png\",\"bust\":\"https://s19.postimg.org/rf8g09uw3/iq.png\"}}},{\"stats\":{\"played\":10,\"wins\":8,\"losses\":2,\"kills\":8,\"deaths\":9,\"playtime\":2103,\"specials\":{\"operatorpvp_blitz_flashedenemy\":\"6\",\"operatorpvp_blitz_flashshieldassist\":\"1\",\"operatorpvp_blitz_flashfollowupkills\":\"1\"}},\"operator\":{\"name\":\"Blitz\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/c8bai9dib/blitz.png\",\"badge\":\"https://s19.postimg.org/m4gyvn8o3/blitz.png\",\"bust\":\"https://s19.postimg.org/pv341vmhv/blitz.png\"}}},{\"stats\":{\"played\":6,\"wins\":5,\"losses\":1,\"kills\":2,\"deaths\":1,\"playtime\":1325,\"specials\":{\"operatorpvp_tachanka_turretkill\":\"2\",\"operatorpvp_tachanka_turretdeployed\":\"8\"}},\"operator\":{\"name\":\"Tachanka\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/5gkvfeoir/tachanka.png\",\"badge\":\"https://s19.postimg.org/6j8u36bhf/tachanka.png\",\"bust\":\"https://s19.postimg.org/5we280n1f/tachanka.png\"}}},{\"stats\":{\"played\":13,\"wins\":6,\"losses\":7,\"kills\":9,\"deaths\":10,\"playtime\":2800,\"specials\":{\"operatorpvp_fuze_clusterchargekill\":\"3\"}},\"operator\":{\"name\":\"Fuze\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/ss8z4i2sj/fuze.png\",\"badge\":\"https://s19.postimg.org/bf7vh4m8z/fuze.png\",\"bust\":\"https://s19.postimg.org/jnrq1pqqr/fuze.png\"}}},{\"stats\":{\"played\":88,\"wins\":47,\"losses\":41,\"kills\":107,\"deaths\":59,\"playtime\":17975,\"specials\":{\"operatorpvp_glaz_sniperkill\":\"96\",\"operatorpvp_glaz_sniperpenetrationkill\":\"18\"}},\"operator\":{\"name\":\"Glaz\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/x03rd9483/glaz.png\",\"badge\":\"https://s19.postimg.org/52sq7aj6r/glaz.png\",\"bust\":\"https://s19.postimg.org/izivirs0z/glaz.png\"}}},{\"stats\":{\"played\":110,\"wins\":60,\"losses\":50,\"kills\":102,\"deaths\":61,\"playtime\":22306,\"specials\":{\"operatorpvp_twitch_shockdronekill\":\"9\",\"operatorpvp_twitch_gadgetdestroybyshockdrone\":\"86\"}},\"operator\":{\"name\":\"Twitch\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/wlxdgjdsz/twitch.png\",\"badge\":\"https://s19.postimg.org/j1pfjr8gz/twitch.png\",\"bust\":\"https://s19.postimg.org/58eegk337/twitch.png\"}}},{\"stats\":{\"played\":470,\"wins\":254,\"losses\":216,\"kills\":485,\"deaths\":289,\"playtime\":95484,\"specials\":{\"operatorpvp_hibana_detonate_projectile\":\"2607\"}},\"operator\":{\"name\":\"Hibana\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/hvgu77usj/hibana.png\",\"bust\":\"https://s19.postimg.org/ma78inhyb/hibana.png\"}}},{\"stats\":{\"played\":84,\"wins\":46,\"losses\":38,\"kills\":71,\"deaths\":51,\"playtime\":18247,\"specials\":{\"operatorpvp_valkyrie_camdeployed\":\"269\"}},\"operator\":{\"name\":\"Valkyrie\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/6bl76o0ir/valkyrie.png\",\"bust\":\"https://s19.postimg.org/55pxfza1v/valkyrie.png\"}}},{\"stats\":{\"played\":58,\"wins\":30,\"losses\":28,\"kills\":31,\"deaths\":41,\"playtime\":11701,\"specials\":{\"operatorpvp_montagne_shieldblockdamage\":\"845\"}},\"operator\":{\"name\":\"Montagne\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/6draoa00z/montagne.png\",\"badge\":\"https://s19.postimg.org/rxvm9y9pf/montagne.png\",\"bust\":\"https://s19.postimg.org/8kcp9uuub/montagne.png\"}}},{\"stats\":{\"played\":26,\"wins\":16,\"losses\":10,\"kills\":15,\"deaths\":18,\"playtime\":5275,\"specials\":{\"operatorpvp_castle_kevlarbarricadedeployed\":\"73\"}},\"operator\":{\"name\":\"Castle\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/d96ha9o6b/castle.png\",\"badge\":\"https://s19.postimg.org/dq1c7k7mr/castle.png\",\"bust\":\"https://s19.postimg.org/hmrzpeftv/castle.png\"}}},{\"stats\":{\"played\":177,\"wins\":85,\"losses\":92,\"kills\":168,\"deaths\":120,\"playtime\":37218,\"specials\":{\"operatorpvp_thatcher_gadgetdestroywithemp\":\"105\"}},\"operator\":{\"name\":\"Thatcher\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/cv55ao42r/thatcher.png\",\"badge\":\"https://s19.postimg.org/cy7uzui77/thatcher.png\",\"bust\":\"https://s19.postimg.org/kmcohqtgj/thatcher.png\"}}},{\"stats\":{\"played\":225,\"wins\":136,\"losses\":89,\"kills\":176,\"deaths\":139,\"playtime\":46028,\"specials\":{\"operatorpvp_mute_gadgetjammed\":\"315\",\"operatorpvp_mute_jammerdeployed\":\"950\"}},\"operator\":{\"name\":\"Mute\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/vkx4rewtf/mute.png\",\"badge\":\"https://s19.postimg.org/5n7r9zcf7/mute.png\",\"bust\":\"https://s19.postimg.org/igds9i0mb/mute.png\"}}},{\"stats\":{\"played\":231,\"wins\":125,\"losses\":106,\"kills\":303,\"deaths\":138,\"playtime\":48592,\"specials\":{\"operatorpvp_blackbeard_gunshieldblockdamage\":\"391\"}},\"operator\":{\"name\":\"Blackbeard\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/t6ewhuc9v/blackbeard.png\",\"bust\":\"https://s19.postimg.org/4pdwl112b/blackbeard.png\"}}},{\"stats\":{\"played\":6,\"wins\":4,\"losses\":2,\"kills\":3,\"deaths\":5,\"playtime\":1502,\"specials\":{\"operatorpvp_frost_dbno\":\"0\"}},\"operator\":{\"name\":\"Frost\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/hdh3dmftv/frost.png\",\"badge\":\"https://s19.postimg.org/aof5bcjvn/frost.png\",\"bust\":\"https://s19.postimg.org/ynfhbti8z/frost.png\"}}},{\"stats\":{\"played\":11,\"wins\":3,\"losses\":8,\"kills\":8,\"deaths\":9,\"playtime\":2269,\"specials\":{\"operatorpvp_kapkan_boobytrapkill\":\"0\",\"operatorpvp_kapkan_boobytrapdeployed\":\"27\"}},\"operator\":{\"name\":\"Kapkan\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/mfttunzqb/kapkan.png\",\"badge\":\"https://s19.postimg.org/jr3midjmr/kapkan.png\",\"bust\":\"https://s19.postimg.org/7yj500aeb/kapkan.png\"}}},{\"stats\":{\"played\":30,\"wins\":15,\"losses\":15,\"kills\":15,\"deaths\":24,\"playtime\":6363,\"specials\":{\"operatorpvp_thermite_chargekill\":\"0\",\"operatorpvp_thermite_chargedeployed\":\"14\",\"operatorpvp_thermite_reinforcementbreached\":\"11\"}},\"operator\":{\"name\":\"Thermite\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/fhknrlvab/thermite.png\",\"badge\":\"https://s19.postimg.org/mwstmbrmr/thermite.png\",\"bust\":\"https://s19.postimg.org/tjxef3lwj/thermite.png\"}}},{\"stats\":{\"played\":42,\"wins\":18,\"losses\":24,\"kills\":30,\"deaths\":32,\"playtime\":8269,\"specials\":{\"operatorpvp_pulse_heartbeatspot\":\"79\",\"operatorpvp_pulse_heartbeatassist\":\"0\"}},\"operator\":{\"name\":\"Pulse\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/6l9vno4o3/pulse.png\",\"badge\":\"https://s19.postimg.org/b0171uppf/pulse.png\",\"bust\":\"https://s19.postimg.org/5ap3bwrz7/pulse.png\"}}},{\"stats\":{\"played\":44,\"wins\":20,\"losses\":24,\"kills\":46,\"deaths\":30,\"playtime\":8666,\"specials\":{\"operatorpvp_ash_bonfirekill\":\"0\",\"operatorpvp_ash_bonfirewallbreached\":\"37\"}},\"operator\":{\"name\":\"Ash\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/a3lvk25k3/ash.png\",\"badge\":\"https://s19.postimg.org/50496pq5f/ash.png\",\"bust\":\"https://s19.postimg.org/wtbtlka03/ash.png\"}}},{\"stats\":{\"played\":192,\"wins\":105,\"losses\":87,\"kills\":182,\"deaths\":121,\"playtime\":38835,\"specials\":{\"operatorpvp_sledge_hammerhole\":\"458\",\"operatorpvp_sledge_hammerkill\":\"0\"}},\"operator\":{\"name\":\"Sledge\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/4blrcwvqb/sledge.png\",\"badge\":\"https://s19.postimg.org/51ndrm6qr/sledge.png\",\"bust\":\"https://s19.postimg.org/a3coskqsj/sledge.png\"}}}]}"
	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: jsonOperator}

	operatorRepository := new(OperatorRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	operatorRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	expectedOperators := new(domain.Operators)
	json.NewDecoder(strings.NewReader(jsonOperator)).Decode(expectedOperators)

	operatorRepository.HttpHandler.Get(url)

	actualSeasons, err := operatorRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Nil(t, err)

	assert.Equal(t, actualSeasons, expectedOperators)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestOperatorInvalidFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonOperator := "{\"operator_records\":[{\"stats\":{\"played\":19,\"wins\":10,\"losses\":9,\"kills\":21,\"deaths\":11,\"playtime\":3968,\"specials\":{\"operatorpvp_doc_selfrevive\":\"0\",\"operatorpvp_doc_hostagerevive\":\"0\",\"operatorpvp_doc_teammaterevive\":\"1\"}},\"operator\":{\"name\":\"Doc\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/i06csah0j/doc.png\",\"badge\":\"https://s19.postimg.org/e5cm0ktk3/doc.png\",\"bust\":\"https://s19.postimg.org/c6dx8rbtv/doc.png\"}}},{\"stats\":{\"played\":19,\"wins\":9,\"losses\":10,\"kills\":22,\"deaths\":13,\"playtime\":4166,\"specials\":{\"operatorpvp_smoke_poisongaskill\":\"1\"}},\"operator\":{\"name\":\"Smoke\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/u4h0vj4ir/smoke.png\",\"badge\":\"https://s19.postimg.org/89rv4nt0j/smoke.png\",\"bust\":\"https://s19.postimg.org/bpspmbo2r/smoke.png\"}}},{\"stats\":{\"played\":149,\"wins\":70,\"losses\":79,\"kills\":115,\"deaths\":101,\"playtime\":30610,\"specials\":{\"operatorpvp_rook_armorboxdeployed\":\"150\",\"operatorpvp_rook_armortakenourself\":\"149\",\"operatorpvp_rook_armortakenteammate\":\"507\"}},\"operator\":{\"name\":\"Rook\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/of5doynqb/rook.png\",\"badge\":\"https://s19.postimg.org/c3lbdtacj/rook.png\",\"bust\":\"https://s19.postimg.org/i27uw5r4z/rook.png\"}}},{\"stats\":{\"played\":94,\"wins\":56,\"losses\":38,\"kills\":62,\"deaths\":51,\"playtime\":20415,\"specials\":{\"operatorpvp_echo_enemy_sonicburst_affected\":\"41\"}},\"operator\":{\"name\":\"Echo\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/vwstsrycj/echo.png\",\"bust\":\"https://s19.postimg.org/c0085o2oz/echo.png\"}}},{\"stats\":{\"played\":31,\"wins\":17,\"losses\":14,\"kills\":24,\"deaths\":22,\"playtime\":6563,\"specials\":{\"operatorpvp_caveira_interrogations\":\"0\"}},\"operator\":{\"name\":\"Capitão\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/l4qnzxtib/capitao.png\",\"bust\":\"https://s19.postimg.org/g6pnyqroj/capitao.png\"}}},{\"stats\":{\"played\":24,\"wins\":14,\"losses\":10,\"kills\":22,\"deaths\":16,\"playtime\":4603,\"specials\":{\"operatorpvp_capitao_lethaldartkills\":\"0\"}},\"operator\":{\"name\":\"Caveira\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/egu2dca03/caveira.png\",\"bust\":\"https://s19.postimg.org/km3hq9eoz/caveira.png\"}}},{\"stats\":{\"played\":104,\"wins\":58,\"losses\":46,\"kills\":104,\"deaths\":70,\"playtime\":21448,\"specials\":{\"operatorpvp_buck_kill\":\"3\"}},\"operator\":{\"name\":\"Buck\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/nprn9ptib/buck.png\",\"badge\":\"https://s19.postimg.org/4sgma7f6r/buck.png\",\"bust\":\"https://s19.postimg.org/5p47gqg83/buck.png\"}}},{\"stats\":{\"played\":666,\"wins\":382,\"losses\":284,\"kills\":704,\"deaths\":416,\"playtime\":141805,\"specials\":{\"operatorpvp_bandit_batterykill\":\"1\"}},\"operator\":{\"name\":\"Bandit\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/hnk0j852b/bandit.png\",\"badge\":\"https://s19.postimg.org/ygjv94wir/bandit.png\",\"bust\":\"https://s19.postimg.org/adb0kfqn7/bandit.png\"}}},{\"stats\":{\"played\":163,\"wins\":90,\"losses\":73,\"kills\":164,\"deaths\":106,\"playtime\":35049,\"specials\":{\"operatorpvp_jager_gadgetdestroybycatcher\":\"180\"}},\"operator\":{\"name\":\"Jäger\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/5x631uc9v/jager.png\",\"badge\":\"https://s19.postimg.org/56mjnjoo3/jager.png\",\"bust\":\"https://s19.postimg.org/nzbqcz89v/jager.png\"}}},{\"stats\":{\"played\":18,\"wins\":9,\"losses\":9,\"kills\":12,\"deaths\":12,\"playtime\":3870,\"specials\":{\"operatorpvp_iq_gadgetspotbyef\":\"178\"}},\"operator\":{\"name\":\"IQ\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/zb1thfezn/iq.png\",\"badge\":\"https://s19.postimg.org/w36ivv7hf/iq.png\",\"bust\":\"https://s19.postimg.org/rf8g09uw3/iq.png\"}}},{\"stats\":{\"played\":10,\"wins\":8,\"losses\":2,\"kills\":8,\"deaths\":9,\"playtime\":2103,\"specials\":{\"operatorpvp_blitz_flashedenemy\":\"6\",\"operatorpvp_blitz_flashshieldassist\":\"1\",\"operatorpvp_blitz_flashfollowupkills\":\"1\"}},\"operator\":{\"name\":\"Blitz\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/c8bai9dib/blitz.png\",\"badge\":\"https://s19.postimg.org/m4gyvn8o3/blitz.png\",\"bust\":\"https://s19.postimg.org/pv341vmhv/blitz.png\"}}},{\"stats\":{\"played\":6,\"wins\":5,\"losses\":1,\"kills\":2,\"deaths\":1,\"playtime\":1325,\"specials\":{\"operatorpvp_tachanka_turretkill\":\"2\",\"operatorpvp_tachanka_turretdeployed\":\"8\"}},\"operator\":{\"name\":\"Tachanka\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/5gkvfeoir/tachanka.png\",\"badge\":\"https://s19.postimg.org/6j8u36bhf/tachanka.png\",\"bust\":\"https://s19.postimg.org/5we280n1f/tachanka.png\"}}},{\"stats\":{\"played\":13,\"wins\":6,\"losses\":7,\"kills\":9,\"deaths\":10,\"playtime\":2800,\"specials\":{\"operatorpvp_fuze_clusterchargekill\":\"3\"}},\"operator\":{\"name\":\"Fuze\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/ss8z4i2sj/fuze.png\",\"badge\":\"https://s19.postimg.org/bf7vh4m8z/fuze.png\",\"bust\":\"https://s19.postimg.org/jnrq1pqqr/fuze.png\"}}},{\"stats\":{\"played\":88,\"wins\":47,\"losses\":41,\"kills\":107,\"deaths\":59,\"playtime\":17975,\"specials\":{\"operatorpvp_glaz_sniperkill\":\"96\",\"operatorpvp_glaz_sniperpenetrationkill\":\"18\"}},\"operator\":{\"name\":\"Glaz\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/x03rd9483/glaz.png\",\"badge\":\"https://s19.postimg.org/52sq7aj6r/glaz.png\",\"bust\":\"https://s19.postimg.org/izivirs0z/glaz.png\"}}},{\"stats\":{\"played\":110,\"wins\":60,\"losses\":50,\"kills\":102,\"deaths\":61,\"playtime\":22306,\"specials\":{\"operatorpvp_twitch_shockdronekill\":\"9\",\"operatorpvp_twitch_gadgetdestroybyshockdrone\":\"86\"}},\"operator\":{\"name\":\"Twitch\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/wlxdgjdsz/twitch.png\",\"badge\":\"https://s19.postimg.org/j1pfjr8gz/twitch.png\",\"bust\":\"https://s19.postimg.org/58eegk337/twitch.png\"}}},{\"stats\":{\"played\":470,\"wins\":254,\"losses\":216,\"kills\":485,\"deaths\":289,\"playtime\":95484,\"specials\":{\"operatorpvp_hibana_detonate_projectile\":\"2607\"}},\"operator\":{\"name\":\"Hibana\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/hvgu77usj/hibana.png\",\"bust\":\"https://s19.postimg.org/ma78inhyb/hibana.png\"}}},{\"stats\":{\"played\":84,\"wins\":46,\"losses\":38,\"kills\":71,\"deaths\":51,\"playtime\":18247,\"specials\":{\"operatorpvp_valkyrie_camdeployed\":\"269\"}},\"operator\":{\"name\":\"Valkyrie\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/6bl76o0ir/valkyrie.png\",\"bust\":\"https://s19.postimg.org/55pxfza1v/valkyrie.png\"}}},{\"stats\":{\"played\":58,\"wins\":30,\"losses\":28,\"kills\":31,\"deaths\":41,\"playtime\":11701,\"specials\":{\"operatorpvp_montagne_shieldblockdamage\":\"845\"}},\"operator\":{\"name\":\"Montagne\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/6draoa00z/montagne.png\",\"badge\":\"https://s19.postimg.org/rxvm9y9pf/montagne.png\",\"bust\":\"https://s19.postimg.org/8kcp9uuub/montagne.png\"}}},{\"stats\":{\"played\":26,\"wins\":16,\"losses\":10,\"kills\":15,\"deaths\":18,\"playtime\":5275,\"specials\":{\"operatorpvp_castle_kevlarbarricadedeployed\":\"73\"}},\"operator\":{\"name\":\"Castle\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/d96ha9o6b/castle.png\",\"badge\":\"https://s19.postimg.org/dq1c7k7mr/castle.png\",\"bust\":\"https://s19.postimg.org/hmrzpeftv/castle.png\"}}},{\"stats\":{\"played\":177,\"wins\":85,\"losses\":92,\"kills\":168,\"deaths\":120,\"playtime\":37218,\"specials\":{\"operatorpvp_thatcher_gadgetdestroywithemp\":\"105\"}},\"operator\":{\"name\":\"Thatcher\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/cv55ao42r/thatcher.png\",\"badge\":\"https://s19.postimg.org/cy7uzui77/thatcher.png\",\"bust\":\"https://s19.postimg.org/kmcohqtgj/thatcher.png\"}}},{\"stats\":{\"played\":225,\"wins\":136,\"losses\":89,\"kills\":176,\"deaths\":139,\"playtime\":46028,\"specials\":{\"operatorpvp_mute_gadgetjammed\":\"315\",\"operatorpvp_mute_jammerdeployed\":\"950\"}},\"operator\":{\"name\":\"Mute\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/vkx4rewtf/mute.png\",\"badge\":\"https://s19.postimg.org/5n7r9zcf7/mute.png\",\"bust\":\"https://s19.postimg.org/igds9i0mb/mute.png\"}}},{\"stats\":{\"played\":231,\"wins\":125,\"losses\":106,\"kills\":303,\"deaths\":138,\"playtime\":48592,\"specials\":{\"operatorpvp_blackbeard_gunshieldblockdamage\":\"391\"}},\"operator\":{\"name\":\"Blackbeard\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/t6ewhuc9v/blackbeard.png\",\"bust\":\"https://s19.postimg.org/4pdwl112b/blackbeard.png\"}}},{\"stats\":{\"played\":6,\"wins\":4,\"losses\":2,\"kills\":3,\"deaths\":5,\"playtime\":1502,\"specials\":{\"operatorpvp_frost_dbno\":\"0\"}},\"operator\":{\"name\":\"Frost\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/hdh3dmftv/frost.png\",\"badge\":\"https://s19.postimg.org/aof5bcjvn/frost.png\",\"bust\":\"https://s19.postimg.org/ynfhbti8z/frost.png\"}}},{\"stats\":{\"played\":11,\"wins\":3,\"losses\":8,\"kills\":8,\"deaths\":9,\"playtime\":2269,\"specials\":{\"operatorpvp_kapkan_boobytrapkill\":\"0\",\"operatorpvp_kapkan_boobytrapdeployed\":\"27\"}},\"operator\":{\"name\":\"Kapkan\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/mfttunzqb/kapkan.png\",\"badge\":\"https://s19.postimg.org/jr3midjmr/kapkan.png\",\"bust\":\"https://s19.postimg.org/7yj500aeb/kapkan.png\"}}},{\"stats\":{\"played\":30,\"wins\":15,\"losses\":15,\"kills\":15,\"deaths\":24,\"playtime\":6363,\"specials\":{\"operatorpvp_thermite_chargekill\":\"0\",\"operatorpvp_thermite_chargedeployed\":\"14\",\"operatorpvp_thermite_reinforcementbreached\":\"11\"}},\"operator\":{\"name\":\"Thermite\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/fhknrlvab/thermite.png\",\"badge\":\"https://s19.postimg.org/mwstmbrmr/thermite.png\",\"bust\":\"https://s19.postimg.org/tjxef3lwj/thermite.png\"}}},{\"stats\":{\"played\":42,\"wins\":18,\"losses\":24,\"kills\":30,\"deaths\":32,\"playtime\":8269,\"specials\":{\"operatorpvp_pulse_heartbeatspot\":\"79\",\"operatorpvp_pulse_heartbeatassist\":\"0\"}},\"operator\":{\"name\":\"Pulse\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/6l9vno4o3/pulse.png\",\"badge\":\"https://s19.postimg.org/b0171uppf/pulse.png\",\"bust\":\"https://s19.postimg.org/5ap3bwrz7/pulse.png\"}}},{\"stats\":{\"played\":44,\"wins\":20,\"losses\":24,\"kills\":46,\"deaths\":30,\"playtime\":8666,\"specials\":{\"operatorpvp_ash_bonfirekill\":\"0\",\"operatorpvp_ash_bonfirewallbreached\":\"37\"}},\"operator\":{\"name\":\"Ash\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/a3lvk25k3/ash.png\",\"badge\":\"https://s19.postimg.org/50496pq5f/ash.png\",\"bust\":\"https://s19.postimg.org/wtbtlka03/ash.png\"}}},{\"stats\":{\"played\":192,\"wins\":105,\"losses\":87,\"kills\":182,\"deaths\":121,\"playtime\":38835,\"specials\":{\"operatorpvp_sledge_hammerhole\":\"458\",\"operatorpvp_sledge_hammerkill\":\"0\"}},\"operator\":{\"name\":\"Sledge\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/4blrcwvqb/sledge.png\",\"badge\":\"https://s19.postimg.org/51ndrm6qr/sledge.png\",\"bust\":\"https://s19.postimg.org/a3coskqsj/sledge.png\"}}}]}"
	mockedHttpContent := HttpContent{StatusCode: 200, Status: "200 OK", Content: jsonOperator}

	operatorRepository := new(OperatorRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	operatorRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, errors.New("Mocked error"))

	operatorRepository.HttpHandler.Get(url)

	_, err := operatorRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}

func TestOperatorValidButWithInvalidHttpRequestFindByProfileNameAndPlatform(t *testing.T) {

	accountName := "accountName"
	platform := "platform"
	unformatedUrl := "https://api.r6stats.com/api/v1/players/%v/operators?platform=%v"

	url := fmt.Sprintf(unformatedUrl, accountName, platform)

	jsonOperator := "{\"operator_records\":[{\"stats\":{\"played\":19,\"wins\":10,\"losses\":9,\"kills\":21,\"deaths\":11,\"playtime\":3968,\"specials\":{\"operatorpvp_doc_selfrevive\":\"0\",\"operatorpvp_doc_hostagerevive\":\"0\",\"operatorpvp_doc_teammaterevive\":\"1\"}},\"operator\":{\"name\":\"Doc\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/i06csah0j/doc.png\",\"badge\":\"https://s19.postimg.org/e5cm0ktk3/doc.png\",\"bust\":\"https://s19.postimg.org/c6dx8rbtv/doc.png\"}}},{\"stats\":{\"played\":19,\"wins\":9,\"losses\":10,\"kills\":22,\"deaths\":13,\"playtime\":4166,\"specials\":{\"operatorpvp_smoke_poisongaskill\":\"1\"}},\"operator\":{\"name\":\"Smoke\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/u4h0vj4ir/smoke.png\",\"badge\":\"https://s19.postimg.org/89rv4nt0j/smoke.png\",\"bust\":\"https://s19.postimg.org/bpspmbo2r/smoke.png\"}}},{\"stats\":{\"played\":149,\"wins\":70,\"losses\":79,\"kills\":115,\"deaths\":101,\"playtime\":30610,\"specials\":{\"operatorpvp_rook_armorboxdeployed\":\"150\",\"operatorpvp_rook_armortakenourself\":\"149\",\"operatorpvp_rook_armortakenteammate\":\"507\"}},\"operator\":{\"name\":\"Rook\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/of5doynqb/rook.png\",\"badge\":\"https://s19.postimg.org/c3lbdtacj/rook.png\",\"bust\":\"https://s19.postimg.org/i27uw5r4z/rook.png\"}}},{\"stats\":{\"played\":94,\"wins\":56,\"losses\":38,\"kills\":62,\"deaths\":51,\"playtime\":20415,\"specials\":{\"operatorpvp_echo_enemy_sonicburst_affected\":\"41\"}},\"operator\":{\"name\":\"Echo\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/vwstsrycj/echo.png\",\"bust\":\"https://s19.postimg.org/c0085o2oz/echo.png\"}}},{\"stats\":{\"played\":31,\"wins\":17,\"losses\":14,\"kills\":24,\"deaths\":22,\"playtime\":6563,\"specials\":{\"operatorpvp_caveira_interrogations\":\"0\"}},\"operator\":{\"name\":\"Capitão\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/l4qnzxtib/capitao.png\",\"bust\":\"https://s19.postimg.org/g6pnyqroj/capitao.png\"}}},{\"stats\":{\"played\":24,\"wins\":14,\"losses\":10,\"kills\":22,\"deaths\":16,\"playtime\":4603,\"specials\":{\"operatorpvp_capitao_lethaldartkills\":\"0\"}},\"operator\":{\"name\":\"Caveira\",\"ctu\":\"BOPE\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/egu2dca03/caveira.png\",\"bust\":\"https://s19.postimg.org/km3hq9eoz/caveira.png\"}}},{\"stats\":{\"played\":104,\"wins\":58,\"losses\":46,\"kills\":104,\"deaths\":70,\"playtime\":21448,\"specials\":{\"operatorpvp_buck_kill\":\"3\"}},\"operator\":{\"name\":\"Buck\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/nprn9ptib/buck.png\",\"badge\":\"https://s19.postimg.org/4sgma7f6r/buck.png\",\"bust\":\"https://s19.postimg.org/5p47gqg83/buck.png\"}}},{\"stats\":{\"played\":666,\"wins\":382,\"losses\":284,\"kills\":704,\"deaths\":416,\"playtime\":141805,\"specials\":{\"operatorpvp_bandit_batterykill\":\"1\"}},\"operator\":{\"name\":\"Bandit\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/hnk0j852b/bandit.png\",\"badge\":\"https://s19.postimg.org/ygjv94wir/bandit.png\",\"bust\":\"https://s19.postimg.org/adb0kfqn7/bandit.png\"}}},{\"stats\":{\"played\":163,\"wins\":90,\"losses\":73,\"kills\":164,\"deaths\":106,\"playtime\":35049,\"specials\":{\"operatorpvp_jager_gadgetdestroybycatcher\":\"180\"}},\"operator\":{\"name\":\"Jäger\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/5x631uc9v/jager.png\",\"badge\":\"https://s19.postimg.org/56mjnjoo3/jager.png\",\"bust\":\"https://s19.postimg.org/nzbqcz89v/jager.png\"}}},{\"stats\":{\"played\":18,\"wins\":9,\"losses\":9,\"kills\":12,\"deaths\":12,\"playtime\":3870,\"specials\":{\"operatorpvp_iq_gadgetspotbyef\":\"178\"}},\"operator\":{\"name\":\"IQ\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/zb1thfezn/iq.png\",\"badge\":\"https://s19.postimg.org/w36ivv7hf/iq.png\",\"bust\":\"https://s19.postimg.org/rf8g09uw3/iq.png\"}}},{\"stats\":{\"played\":10,\"wins\":8,\"losses\":2,\"kills\":8,\"deaths\":9,\"playtime\":2103,\"specials\":{\"operatorpvp_blitz_flashedenemy\":\"6\",\"operatorpvp_blitz_flashshieldassist\":\"1\",\"operatorpvp_blitz_flashfollowupkills\":\"1\"}},\"operator\":{\"name\":\"Blitz\",\"ctu\":\"GSG 9\",\"images\":{\"figure\":\"https://s19.postimg.org/c8bai9dib/blitz.png\",\"badge\":\"https://s19.postimg.org/m4gyvn8o3/blitz.png\",\"bust\":\"https://s19.postimg.org/pv341vmhv/blitz.png\"}}},{\"stats\":{\"played\":6,\"wins\":5,\"losses\":1,\"kills\":2,\"deaths\":1,\"playtime\":1325,\"specials\":{\"operatorpvp_tachanka_turretkill\":\"2\",\"operatorpvp_tachanka_turretdeployed\":\"8\"}},\"operator\":{\"name\":\"Tachanka\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/5gkvfeoir/tachanka.png\",\"badge\":\"https://s19.postimg.org/6j8u36bhf/tachanka.png\",\"bust\":\"https://s19.postimg.org/5we280n1f/tachanka.png\"}}},{\"stats\":{\"played\":13,\"wins\":6,\"losses\":7,\"kills\":9,\"deaths\":10,\"playtime\":2800,\"specials\":{\"operatorpvp_fuze_clusterchargekill\":\"3\"}},\"operator\":{\"name\":\"Fuze\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/ss8z4i2sj/fuze.png\",\"badge\":\"https://s19.postimg.org/bf7vh4m8z/fuze.png\",\"bust\":\"https://s19.postimg.org/jnrq1pqqr/fuze.png\"}}},{\"stats\":{\"played\":88,\"wins\":47,\"losses\":41,\"kills\":107,\"deaths\":59,\"playtime\":17975,\"specials\":{\"operatorpvp_glaz_sniperkill\":\"96\",\"operatorpvp_glaz_sniperpenetrationkill\":\"18\"}},\"operator\":{\"name\":\"Glaz\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/x03rd9483/glaz.png\",\"badge\":\"https://s19.postimg.org/52sq7aj6r/glaz.png\",\"bust\":\"https://s19.postimg.org/izivirs0z/glaz.png\"}}},{\"stats\":{\"played\":110,\"wins\":60,\"losses\":50,\"kills\":102,\"deaths\":61,\"playtime\":22306,\"specials\":{\"operatorpvp_twitch_shockdronekill\":\"9\",\"operatorpvp_twitch_gadgetdestroybyshockdrone\":\"86\"}},\"operator\":{\"name\":\"Twitch\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/wlxdgjdsz/twitch.png\",\"badge\":\"https://s19.postimg.org/j1pfjr8gz/twitch.png\",\"bust\":\"https://s19.postimg.org/58eegk337/twitch.png\"}}},{\"stats\":{\"played\":470,\"wins\":254,\"losses\":216,\"kills\":485,\"deaths\":289,\"playtime\":95484,\"specials\":{\"operatorpvp_hibana_detonate_projectile\":\"2607\"}},\"operator\":{\"name\":\"Hibana\",\"ctu\":\"SAT\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/hvgu77usj/hibana.png\",\"bust\":\"https://s19.postimg.org/ma78inhyb/hibana.png\"}}},{\"stats\":{\"played\":84,\"wins\":46,\"losses\":38,\"kills\":71,\"deaths\":51,\"playtime\":18247,\"specials\":{\"operatorpvp_valkyrie_camdeployed\":\"269\"}},\"operator\":{\"name\":\"Valkyrie\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/6bl76o0ir/valkyrie.png\",\"bust\":\"https://s19.postimg.org/55pxfza1v/valkyrie.png\"}}},{\"stats\":{\"played\":58,\"wins\":30,\"losses\":28,\"kills\":31,\"deaths\":41,\"playtime\":11701,\"specials\":{\"operatorpvp_montagne_shieldblockdamage\":\"845\"}},\"operator\":{\"name\":\"Montagne\",\"ctu\":\"GIGN\",\"images\":{\"figure\":\"https://s19.postimg.org/6draoa00z/montagne.png\",\"badge\":\"https://s19.postimg.org/rxvm9y9pf/montagne.png\",\"bust\":\"https://s19.postimg.org/8kcp9uuub/montagne.png\"}}},{\"stats\":{\"played\":26,\"wins\":16,\"losses\":10,\"kills\":15,\"deaths\":18,\"playtime\":5275,\"specials\":{\"operatorpvp_castle_kevlarbarricadedeployed\":\"73\"}},\"operator\":{\"name\":\"Castle\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/d96ha9o6b/castle.png\",\"badge\":\"https://s19.postimg.org/dq1c7k7mr/castle.png\",\"bust\":\"https://s19.postimg.org/hmrzpeftv/castle.png\"}}},{\"stats\":{\"played\":177,\"wins\":85,\"losses\":92,\"kills\":168,\"deaths\":120,\"playtime\":37218,\"specials\":{\"operatorpvp_thatcher_gadgetdestroywithemp\":\"105\"}},\"operator\":{\"name\":\"Thatcher\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/cv55ao42r/thatcher.png\",\"badge\":\"https://s19.postimg.org/cy7uzui77/thatcher.png\",\"bust\":\"https://s19.postimg.org/kmcohqtgj/thatcher.png\"}}},{\"stats\":{\"played\":225,\"wins\":136,\"losses\":89,\"kills\":176,\"deaths\":139,\"playtime\":46028,\"specials\":{\"operatorpvp_mute_gadgetjammed\":\"315\",\"operatorpvp_mute_jammerdeployed\":\"950\"}},\"operator\":{\"name\":\"Mute\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/vkx4rewtf/mute.png\",\"badge\":\"https://s19.postimg.org/5n7r9zcf7/mute.png\",\"bust\":\"https://s19.postimg.org/igds9i0mb/mute.png\"}}},{\"stats\":{\"played\":231,\"wins\":125,\"losses\":106,\"kills\":303,\"deaths\":138,\"playtime\":48592,\"specials\":{\"operatorpvp_blackbeard_gunshieldblockdamage\":\"391\"}},\"operator\":{\"name\":\"Blackbeard\",\"ctu\":\"Navy Seal\",\"images\":{\"figure\":null,\"badge\":\"https://s19.postimg.org/t6ewhuc9v/blackbeard.png\",\"bust\":\"https://s19.postimg.org/4pdwl112b/blackbeard.png\"}}},{\"stats\":{\"played\":6,\"wins\":4,\"losses\":2,\"kills\":3,\"deaths\":5,\"playtime\":1502,\"specials\":{\"operatorpvp_frost_dbno\":\"0\"}},\"operator\":{\"name\":\"Frost\",\"ctu\":\"JTF2\",\"images\":{\"figure\":\"https://s19.postimg.org/hdh3dmftv/frost.png\",\"badge\":\"https://s19.postimg.org/aof5bcjvn/frost.png\",\"bust\":\"https://s19.postimg.org/ynfhbti8z/frost.png\"}}},{\"stats\":{\"played\":11,\"wins\":3,\"losses\":8,\"kills\":8,\"deaths\":9,\"playtime\":2269,\"specials\":{\"operatorpvp_kapkan_boobytrapkill\":\"0\",\"operatorpvp_kapkan_boobytrapdeployed\":\"27\"}},\"operator\":{\"name\":\"Kapkan\",\"ctu\":\"Spetsnaz\",\"images\":{\"figure\":\"https://s19.postimg.org/mfttunzqb/kapkan.png\",\"badge\":\"https://s19.postimg.org/jr3midjmr/kapkan.png\",\"bust\":\"https://s19.postimg.org/7yj500aeb/kapkan.png\"}}},{\"stats\":{\"played\":30,\"wins\":15,\"losses\":15,\"kills\":15,\"deaths\":24,\"playtime\":6363,\"specials\":{\"operatorpvp_thermite_chargekill\":\"0\",\"operatorpvp_thermite_chargedeployed\":\"14\",\"operatorpvp_thermite_reinforcementbreached\":\"11\"}},\"operator\":{\"name\":\"Thermite\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/fhknrlvab/thermite.png\",\"badge\":\"https://s19.postimg.org/mwstmbrmr/thermite.png\",\"bust\":\"https://s19.postimg.org/tjxef3lwj/thermite.png\"}}},{\"stats\":{\"played\":42,\"wins\":18,\"losses\":24,\"kills\":30,\"deaths\":32,\"playtime\":8269,\"specials\":{\"operatorpvp_pulse_heartbeatspot\":\"79\",\"operatorpvp_pulse_heartbeatassist\":\"0\"}},\"operator\":{\"name\":\"Pulse\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/6l9vno4o3/pulse.png\",\"badge\":\"https://s19.postimg.org/b0171uppf/pulse.png\",\"bust\":\"https://s19.postimg.org/5ap3bwrz7/pulse.png\"}}},{\"stats\":{\"played\":44,\"wins\":20,\"losses\":24,\"kills\":46,\"deaths\":30,\"playtime\":8666,\"specials\":{\"operatorpvp_ash_bonfirekill\":\"0\",\"operatorpvp_ash_bonfirewallbreached\":\"37\"}},\"operator\":{\"name\":\"Ash\",\"ctu\":\"FBI SWAT\",\"images\":{\"figure\":\"https://s19.postimg.org/a3lvk25k3/ash.png\",\"badge\":\"https://s19.postimg.org/50496pq5f/ash.png\",\"bust\":\"https://s19.postimg.org/wtbtlka03/ash.png\"}}},{\"stats\":{\"played\":192,\"wins\":105,\"losses\":87,\"kills\":182,\"deaths\":121,\"playtime\":38835,\"specials\":{\"operatorpvp_sledge_hammerhole\":\"458\",\"operatorpvp_sledge_hammerkill\":\"0\"}},\"operator\":{\"name\":\"Sledge\",\"ctu\":\"SAS\",\"images\":{\"figure\":\"https://s19.postimg.org/4blrcwvqb/sledge.png\",\"badge\":\"https://s19.postimg.org/51ndrm6qr/sledge.png\",\"bust\":\"https://s19.postimg.org/a3coskqsj/sledge.png\"}}}]}"
	mockedHttpContent := HttpContent{StatusCode: 400, Status: "400 BAD REQUEST", Content: jsonOperator}

	operatorRepository := new(OperatorRepository)
	mockedRepositoryHttpHandler := new(MockedRepositoryHttpHandler)

	operatorRepository.HttpHandler = mockedRepositoryHttpHandler

	mockedRepositoryHttpHandler.On("Get", url).Return(mockedHttpContent, nil)

	operatorRepository.HttpHandler.Get(url)

	_, err := operatorRepository.FindByProfileNameAndPlatform(accountName, platform)

	assert.Error(t, err)

	mockedRepositoryHttpHandler.AssertExpectations(t)

}
