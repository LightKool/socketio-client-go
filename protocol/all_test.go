package protocol

import (
	"encoding/json"
	"fmt"
	"testing"
)

var data = `42/room,13["getScreenData",{"data":{"dictionary":{"turnover":"交易咨询额","register_users":"注册用户数","annualized_return_rate":"年化收益率","profit_rate_bank":"银行平均收益率","profit_rate_platform":"酷邀贷各业务渠道平均收益率","bad_debt_rate":"坏账率","bid_failure_rate":"流标率","finish_rate":"投标进度","invest_period":"投资周期","start_money":"起投金额","area_name":"地域名称","target_name":"标的名称","type":"类型","name":"名称","date":"日期","datetime":"时间","phone_number":"手机号","money":"交易金额","hour":"小时","pv":"页面浏览量","uv":"用户浏览量","shop_distribution":"门店分布","invest_amount":"投资额","order_amount":"交易频次","shop_name":"店铺名称","transaction_amount":"交易金额","m_active":"月活跃","d_active":"日活跃"},"profit_rate":[{"date":"2018-03-02","profit_rate_bank":"0.023","profit_rate_platform":"0.061"},{"date":"2018-03-09","profit_rate_bank":"0.02","profit_rate_platform":"0.06"},{"date":"2018-03-16","profit_rate_bank":"0.022","profit_rate_platform":"0.061"},{"date":"2018-03-23","profit_rate_bank":"0.03","profit_rate_platform":"0.062"},{"date":"2018-03-30","profit_rate_bank":"0.022","profit_rate_platform":"0.061"}],"pv":[{"hour":"2","pv":"37"},{"hour":"3","pv":"26"},{"hour":"4","pv":"48"},{"hour":"5","pv":"57"},{"hour":"6","pv":"176"},{"hour":"7","pv":"171"},{"hour":"8","pv":"513"},{"hour":"9","pv":"2567"},{"hour":"10","pv":"3233"},{"hour":"11","pv":"1190"},{"hour":"12","pv":"684"},{"hour":"13","pv":"453"}],"total":{"card_bound":"317972","investors":"250541","is_insert":"0","register_users":4627818,"register_users_base":"0","repurchase":"31283","turnover":"40221099566","turnover_base":"0","update_timestamp":"1522388734000"},"turnover_area":[{"area_name":"上海","register_users":"9016","turnover":"9909855917.00"},{"area_name":"云南","register_users":"3260","turnover":"45251500.00"},{"area_name":"内蒙古","register_users":"4177","turnover":"95065902.00"},{"area_name":"北京","register_users":"503","turnover":"409811777.00"},{"area_name":"吉林","register_users":"4174","turnover":"239140800.00"},{"area_name":"四川","register_users":"11693","turnover":"331752800.00"},{"area_name":"天津","register_users":"945","turnover":"144021000.00"},{"area_name":"宁夏","register_users":"585","turnover":"28629500.00"},{"area_name":"安徽","register_users":"9436","turnover":"709242500.00"},{"area_name":"山东","register_users":"29606","turnover":"1902576508.00"},{"area_name":"山西","register_users":"15494","turnover":"325560200.00"},{"area_name":"广东","register_users":"11314","turnover":"1041614339.00"},{"area_name":"广西","register_users":"6838","turnover":"110255600.00"},{"area_name":"新疆","register_users":"825","turnover":"102125500.00"},{"area_name":"江苏","register_users":"14845","turnover":"2747900375.00"},{"area_name":"江西","register_users":"8159","turnover":"388737800.00"},{"area_name":"河北","register_users":"18068","turnover":"525001972.00"},{"area_name":"河南","register_users":"32239","turnover":"640874178.00"},{"area_name":"浙江","register_users":"8520","turnover":"2262975710.00"},{"area_name":"海南","register_users":"694","turnover":"8858500.00"},{"area_name":"港澳","register_users":"43","turnover":"168462300.00"},{"area_name":"湖北","register_users":"11587","turnover":"355408100.00"},{"area_name":"湖南","register_users":"14026","turnover":"578115000.00"},{"area_name":"甘肃","register_users":"2442","turnover":"123741300.00"},{"area_name":"福建","register_users":"7437","turnover":"476930000.00"},{"area_name":"西藏","register_users":"11","turnover":"2036000.00"},{"area_name":"贵州","register_users":"4847","turnover":"121431100.00"},{"area_name":"辽宁","register_users":"5261","turnover":"252577800.00"},{"area_name":"重庆","register_users":"3329","turnover":"47062500.00"},{"area_name":"陕西","register_users":"4100","turnover":"232593100.00"},{"area_name":"青海","register_users":"271","turnover":"53754100.00"},{"area_name":"黑龙江","register_users":"6785","turnover":"264716677.00"}],"turnover_today":[{"hour":"2","turnover":"0"},{"hour":"3","turnover":"0"},{"hour":"4","turnover":"0"},{"hour":"5","turnover":"0"},{"hour":"6","turnover":"0"},{"hour":"7","turnover":"550000.00"},{"hour":"8","turnover":"0"},{"hour":"9","turnover":"6790000.00"},{"hour":"10","turnover":"12860000.00"},{"hour":"11","turnover":"2830000.00"},{"hour":"12","turnover":"700000.00"},{"hour":"13","turnover":"890000.00"}],"user_register_trend":[{"datetime":"2018-03-30 13:16:29","id":"0","mobile":"150****0178","username":"c**"},{"datetime":"2018-03-30 12:49:23","id":"1","mobile":"180****6665","username":"h**"},{"datetime":"2018-03-30 12:39:53","id":"2","mobile":"182****2028","username":"x**"},{"datetime":"2018-03-30 12:01:07","id":"3","mobile":"188****2525","username":"l**"},{"datetime":"2018-03-30 10:56:52","id":"4","mobile":"137****7823","username":"m**"},{"datetime":"2018-03-30 10:55:03","id":"5","mobile":"158****0517","username":"w**"},{"datetime":"2018-03-30 10:51:24","id":"6","mobile":"133****8388","username":"w**"},{"datetime":"2018-03-30 10:51:01","id":"7","mobile":"139****2495","username":"w**"},{"datetime":"2018-03-30 10:50:45","id":"8","mobile":"158****4251","username":"m**"},{"datetime":"2018-03-30 10:49:28","id":"9","mobile":"137****8113","username":"l**"},{"datetime":"2018-03-30 10:48:58","id":"10","mobile":"138****9767","username":"z**"},{"datetime":"2018-03-30 10:48:27","id":"11","mobile":"189****1833","username":"c**"},{"datetime":"2018-03-30 10:40:10","id":"12","mobile":"138****1219","username":"a**"},{"datetime":"2018-03-30 10:35:54","id":"13","mobile":"189****5811","username":"x**"},{"datetime":"2018-03-30 10:29:39","id":"14","mobile":"136****2087","username":"l**"},{"datetime":"2018-03-30 10:23:44","id":"15","mobile":"180****9789","username":"D**"},{"datetime":"2018-03-30 10:14:32","id":"16","mobile":"180****6339","username":"x**"},{"datetime":"2018-03-30 10:02:09","id":"17","mobile":"136****2070","username":"y**"},{"datetime":"2018-03-30 09:47:07","id":"18","mobile":"132****0951","username":"s**"},{"datetime":"2018-03-30 09:45:09","id":"19","mobile":"157****8841","username":"g**"}],"user_transaction_trend":[{"area_name":"上海市","datetime":"2018-03-30 13:26:14","id":"0","money":"310000.00","phone_number":"189****1045","type":"交易咨询","username":"朱**"},{"area_name":"上海市","datetime":"2018-03-30 13:24:44","id":"1","money":"310000.00","phone_number":"189****1045","type":"交易咨询","username":"朱**"},{"area_name":"浙江省","datetime":"2018-03-30 13:17:08","id":"2","money":"50000.00","phone_number":"159****4570","type":"交易咨询","username":"葛**"},{"area_name":"上海市","datetime":"2018-03-30 13:11:29","id":"3","money":"170000.00","phone_number":"135****3955","type":"交易咨询","username":"王**"},{"area_name":"江苏省","datetime":"2018-03-30 13:09:21","id":"4","money":"50000.00","phone_number":"137****5854","type":"交易咨询","username":"江**"},{"area_name":"上海市","datetime":"2018-03-30 12:51:50","id":"5","money":"100000.00","phone_number":"136****9681","type":"交易咨询","username":"俞**"},{"area_name":"安徽省","datetime":"2018-03-30 12:32:21","id":"6","money":"200000.00","phone_number":"133****6665","type":"交易咨询","username":"张**"},{"area_name":"上海市","datetime":"2018-03-30 12:26:08","id":"7","money":"50000.00","phone_number":"130****9886","type":"交易咨询","username":"包**"},{"area_name":"上海市","datetime":"2018-03-30 12:01:09","id":"8","money":"350000.00","phone_number":"138****6297","type":"交易咨询","username":"朱**"},{"area_name":"江苏省","datetime":"2018-03-30 11:56:20","id":"9","money":"100000.00","phone_number":"138****0330","type":"交易咨询","username":"张**"},{"area_name":"山东省","datetime":"2018-03-30 11:55:55","id":"10","money":"50000.00","phone_number":"150****5264","type":"交易咨询","username":"尚**"},{"area_name":"上海市","datetime":"2018-03-30 11:54:46","id":"11","money":"150000.00","phone_number":"134****7054","type":"交易咨询","username":"王**"},{"area_name":"上海市","datetime":"2018-03-30 11:49:32","id":"12","money":"100000.00","phone_number":"133****3557","type":"交易咨询","username":"王**"},{"area_name":"安徽省","datetime":"2018-03-30 11:48:54","id":"13","money":"50000.00","phone_number":"159****8476","type":"交易咨询","username":"胡**"},{"area_name":"上海市","datetime":"2018-03-30 11:44:10","id":"14","money":"50000.00","phone_number":"135****7115","type":"交易咨询","username":"钱**"},{"area_name":"甘肃省","datetime":"2018-03-30 11:42:38","id":"15","money":"50000.00","phone_number":"139****4320","type":"交易咨询","username":"唐**"},{"area_name":"湖南省","datetime":"2018-03-30 11:40:06","id":"16","money":"200000.00","phone_number":"183****8091","type":"交易咨询","username":"杜**"},{"area_name":"上海市","datetime":"2018-03-30 11:39:44","id":"17","money":"50000.00","phone_number":"133****6776","type":"交易咨询","username":"汪**"},{"area_name":"浙江省","datetime":"2018-03-30 11:33:44","id":"18","money":"50000.00","phone_number":"137****4047","type":"交易咨询","username":"汪**"},{"area_name":"上海市","datetime":"2018-03-30 11:31:53","id":"19","money":"50000.00","phone_number":"135****0280","type":"交易咨询","username":"章**"}],"uv":[{"hour":"2","uv":"20"},{"hour":"3","uv":"15"},{"hour":"4","uv":"28"},{"hour":"5","uv":"30"},{"hour":"6","uv":"60"},{"hour":"7","uv":"57"},{"hour":"8","uv":"138"},{"hour":"9","uv":"423"},{"hour":"10","uv":"490"},{"hour":"11","uv":"242"},{"hour":"12","uv":"172"},{"hour":"13","uv":"123"}],"m_active":[{"month":"Apr-17","m_active":"60656"},{"month":"May-17","m_active":"97461"},{"month":"Jun-17","m_active":"227523"},{"month":"Jul-17","m_active":"332858"},{"month":"Aug-17","m_active":"225080"},{"month":"Sep-17","m_active":"121800"}],"d_active":[{"day":"2018-03-20","d_active":2138},{"day":"2018-03-21","d_active":2070},{"day":"2018-03-22","d_active":1909},{"day":"2018-03-23","d_active":1652},{"day":"2018-03-24","d_active":947},{"day":"2018-03-25","d_active":912},{"day":"2018-03-26","d_active":2000},{"day":"2018-03-27","d_active":2352},{"day":"2018-03-28","d_active":3061},{"day":"2018-03-29","d_active":2488}]},"code":0,"msg":"Succeed"}]`

func TestDecodePacket(t *testing.T) {
	p, err := decodePacket(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p.payload)
}

func TestPacketDecodeMessage(t *testing.T) {
	p, err := decodePacket(data)
	if err != nil {
		t.Fatal(err)
	}
	m, err := p.DecodeMessage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m.Type, m.Namespace, m.ID, m.Payloads)
}

func TestJSON(t *testing.T) {
	data := `["m","123456"]`
	var value []interface{}
	err := json.Unmarshal([]byte(data), &value)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(value)
}

func TestMessageEncodePacket(t *testing.T) {
	m := &Message{
		Type:      MessageTypeEvent,
		Namespace: "/room",
		ID:        1,
		Payloads: []interface{}{
			"SubAdd",
			[]string{"sub1", "sub2"},
		},
	}
	p, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p.String())
}

func TestWebSocket(t *testing.T) {
	url := "wss://streamer.cryptocompare.com/socket.io/?EIO=3&transport=websocket"
	tr := NewWebSocketTransport()
	conn, err := tr.Dial(url)
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	for index := 0; index < 2; index++ {
		p, err := conn.Read()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
	}
	m := &Message{
		Type:  MessageTypeEvent,
		ID:    -1,
		Event: "SubAdd",
		Payloads: []interface{}{
			map[string]interface{}{
				"subs": []string{"0~Binance~BTC~USDT"},
			},
		},
	}
	p, _ := m.Encode()
	err = conn.Write(p)
	fmt.Println(p)
	if err != nil {
		t.Fatal(err)
	}
	for index := 0; index < 10; index++ {
		p, err = conn.Read()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(p)
	}
}
