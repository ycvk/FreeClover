package openShamrockEntity

// ////////////////////////////////////////////////////////////账户

type Common struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Msg     string `json:"msg"`
	Wording string `json:"wording,omitempty"`
	Echo    string `json:"echo,omitempty"`
}

type LoginInfo struct {
	Common
	Data struct {
		UserId   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"data,omitempty"`
}
type PhoneModel struct {
	Common
	Data struct {
		Variants []struct {
			ModelShow string `json:"model_show"`
			NeedPay   bool   `json:"need_pay"`
		} `json:"variants"`
	} `json:"data,omitempty"`
}
type Device struct {
	Common
	Data struct {
		AppID      int64  `json:"app_id"`
		DeviceName string `json:"device_name"`
		DeviceKind string `json:"device_kind"`
	} `json:"data,omitempty"`
}

// ////////////////////////////////////////////////////////////联系人

type StrangerInfo struct {
	Common
	Data struct {
		UserId    string      `json:"user_id"`
		Nickname  string      `json:"nickname"`
		Age       int         `json:"age"`
		Sex       string      `json:"sex"`
		Level     int         `json:"level"`
		LoginDays int         `json:"login_days"`
		Qid       string      `json:"qid"`
		Vote      int         `json:"vote"`
		WzryHonor interface{} `json:"wzry_honor"`
		Ext       struct {
			AddSrcId                   int         `json:"add_src_id"`
			AddSrcName                 string      `json:"add_src_name"`
			AddSubSrcId                int         `json:"add_sub_src_id"`
			AllowCalInteractive        bool        `json:"allow_cal_interactive"`
			AllowClick                 bool        `json:"allow_click"`
			AllowPeopleSee             bool        `json:"allow_people_see"`
			AuthState                  int         `json:"auth_state"`
			BigClubVipOpen             int         `json:"big_club_vip_open"`
			HollywoodVipOpen           int         `json:"hollywood_vip_open"`
			QqVipOpen                  int         `json:"qq_vip_open"`
			SuperQqOpen                int         `json:"super_qq_open"`
			SuperVipOpen               int         `json:"super_vip_open"`
			Voted                      int         `json:"voted"`
			BabyQSwitch                bool        `json:"baby_q_switch"`
			BindPhoneInfo              string      `json:"bind_phone_info"`
			CardId                     int         `json:"card_id"`
			CardType                   int         `json:"card_type"`
			Category                   int         `json:"category"`
			ClothesId                  int         `json:"clothes_id"`
			CoverUrl                   string      `json:"cover_url"`
			Declaration                interface{} `json:"declaration"`
			DefaultCardId              int         `json:"default_card_id"`
			DiyComplicatedInfo         interface{} `json:"diy_complicated_info"`
			DiyDefaultText             interface{} `json:"diy_default_text"`
			DiyText                    interface{} `json:"diy_text"`
			DiyTextDegree              float64     `json:"diy_text_degree"`
			DiyTextFontId              int         `json:"diy_text_font_id"`
			DiyTextHeight              float64     `json:"diy_text_height"`
			DiyTextWidth               float64     `json:"diy_text_width"`
			DiyTextLocX                float64     `json:"diy_text_loc_x"`
			DiyTextLocY                float64     `json:"diy_text_loc_y"`
			DressUpIsOn                bool        `json:"dress_up_is_on"`
			EncId                      interface{} `json:"enc_id"`
			EnlargeQzonePic            int         `json:"enlarge_qzone_pic"`
			ExtendFriendEntryAddFriend int         `json:"extend_friend_entry_add_friend"`
			ExtendFriendEntryContact   int         `json:"extend_friend_entry_contact"`
			ExtendFriendFlag           int         `json:"extend_friend_flag"`
			ExtendFriendQuestion       int         `json:"extend_friend_question"`
			ExtendFriendVoiceDuration  int         `json:"extend_friend_voice_duration"`
			FavoriteSource             int         `json:"favorite_source"`
			FeedPreviewTime            int         `json:"feed_preview_time"`
			FontId                     int         `json:"font_id"`
			FontType                   int         `json:"font_type"`
			QidBgUrl                   string      `json:"qid_bg_url"`
			QidColor                   string      `json:"qid_color"`
			QidLogoUrl                 string      `json:"qid_logo_url"`
			QqCardIsOn                 bool        `json:"qq_card_is_on"`
			SchoolId                   interface{} `json:"school_id"`
			SchoolName                 interface{} `json:"school_name"`
			SchoolVerifiedFlag         bool        `json:"school_verified_flag"`
			ShowPublishButton          bool        `json:"show_publish_button"`
			Singer                     string      `json:"singer"`
			SongDura                   int         `json:"song_dura"`
			SongId                     string      `json:"song_id"`
			SongName                   string      `json:"song_name"`
		} `json:"ext"`
	} `json:"data,omitempty"`
}
type FriendList struct {
	Common
	Data []struct {
		UserId          int64  `json:"user_id"`
		UserName        string `json:"user_name"`
		UserDisplayname string `json:"user_displayname"`
		UserRemark      string `json:"user_remark"`
		Age             int    `json:"age"`
		Gender          int    `json:"gender"`
		GroupId         int    `json:"group_id"`
		Platform        string `json:"platform"`
		TermType        int    `json:"term_type"`
	} `json:"data,omitempty"`
}
type UnidirectionalFriend struct {
	Common
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Source   string `json:"source"`
}
type GroupInfo struct {
	Common
	Data struct {
		GroupId        int    `json:"group_id"`
		GroupName      string `json:"group_name"`
		GroupRemark    string `json:"group_remark"`
		GroupUin       int    `json:"group_uin"`
		Admins         []int  `json:"admins"`
		ClassText      string `json:"class_text"`
		IsFrozen       bool   `json:"is_frozen"`
		MaxMember      int    `json:"max_member"`
		MemberNum      int    `json:"member_num"`
		MemberCount    int    `json:"member_count"`
		MaxMemberCount int    `json:"max_member_count"`
	} `json:"data,omitempty"`
}
type GroupList struct {
	Common
	Data []struct {
		GroupId        int    `json:"group_id"`
		GroupName      string `json:"group_name"`
		GroupRemark    string `json:"group_remark"`
		GroupUin       int    `json:"group_uin"`
		Admins         []int  `json:"admins"`
		ClassText      string `json:"class_text"`
		IsFrozen       bool   `json:"is_frozen"`
		MaxMember      int    `json:"max_member"`
		MemberNum      int    `json:"member_num"`
		MemberCount    int    `json:"member_count"`
		MaxMemberCount int    `json:"max_member_count"`
	} `json:"data,omitempty"`
}
type GroupMemberInfo struct {
	Common
	Data struct {
		UserId          int    `json:"user_id"`
		GroupId         int    `json:"group_id"`
		UserName        string `json:"user_name"`
		Sex             string `json:"sex"`
		Title           string `json:"title"`
		TitleExpireTime int    `json:"title_expire_time"`
		Nickname        string `json:"nickname"`
		UserDisplayname string `json:"user_displayname"`
		Distance        int    `json:"distance"`
		Honor           []int  `json:"honor"`
		JoinTime        int    `json:"join_time"`
		LastActiveTime  int    `json:"last_active_time"`
		LastSentTime    int    `json:"last_sent_time"`
		UniqueName      string `json:"unique_name"`
		Area            string `json:"area"`
		Level           int    `json:"level"`
		Role            string `json:"role"`
		Unfriendly      bool   `json:"unfriendly"`
		CardChangeable  bool   `json:"card_changeable"`
	} `json:"data,omitempty"`
}

type GroupMemberList struct {
	Common
	Data []struct {
		UserId          int    `json:"user_id"`
		GroupId         int    `json:"group_id"`
		UserName        string `json:"user_name"`
		Sex             string `json:"sex"`
		Title           string `json:"title"`
		TitleExpireTime int    `json:"title_expire_time"`
		Nickname        string `json:"nickname"`
		UserDisplayname string `json:"user_displayname"`
		Distance        int    `json:"distance"`
		Honor           []int  `json:"honor"`
		JoinTime        int    `json:"join_time"`
		LastActiveTime  int    `json:"last_active_time"`
		LastSentTime    int    `json:"last_sent_time"`
		UniqueName      string `json:"unique_name"`
		Area            string `json:"area"`
		Level           int    `json:"level"`
		Role            string `json:"role"`
		Unfriendly      bool   `json:"unfriendly"`
		CardChangeable  bool   `json:"card_changeable"`
	} `json:"data,omitempty"`
}
type HonorInfo struct {
	UserID      int64  `json:"user_id"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	DayCount    int32  `json:"day_count"`
	ID          int32  `json:"id"`
	Description string `json:"description"`
}
type GroupHonorInfo struct {
	GroupID          int64       `json:"group_id"`
	CurrentTalkative []HonorInfo `json:"current_talkative"`
	TalkativeList    []HonorInfo `json:"talkative_list"`
	PerformerList    []HonorInfo `json:"performer_list"`
	LegendList       []HonorInfo `json:"legend_list"`
	StrongNewbieList []HonorInfo `json:"strong_newbie_list"`
	EmotionList      []HonorInfo `json:"emotion_list"`
	All              []HonorInfo `json:"all"`
}
type GroupSystemMsg struct {
	Common
	Data struct {
		InvitedRequests []struct {
			RequestID    int64  `json:"request_id"`
			InvitorUIN   int64  `json:"invitor_uin"`
			InvitorNick  string `json:"invitor_nick"`
			GroupID      int64  `json:"group_id"`
			GroupName    string `json:"group_name"`
			Checked      bool   `json:"checked"`
			Actor        int64  `json:"actor"`
			RequesterUIN int64  `json:"requester_uin"`
			Flag         string `json:"flag"`
		} `json:"invited_requests"`
		JoinRequests []struct {
			RequestID     int64  `json:"request_id"`
			RequesterUIN  int64  `json:"requester_uin"`
			RequesterNick string `json:"requester_nick"`
			Message       string `json:"message"`
			GroupID       int64  `json:"group_id"`
			GroupName     string `json:"group_name"`
			Checked       bool   `json:"checked"`
			Actor         int64  `json:"actor"`
			Flag          string `json:"flag"`
		} `json:"join_requests"`
	} `json:"data,omitempty"`
}
type FriendSystemMsg struct {
	Common
	Data []struct {
		RequestId       int64  `json:"request_id"`
		RequesterUin    int    `json:"requester_uin"`
		RequesterNick   string `json:"requester_nick"`
		Source          string `json:"source"`
		Message         string `json:"message"`
		SourceGroupName string `json:"source_group_name"`
		SourceGroupId   int    `json:"source_group_id"`
		Flag            string `json:"flag"`
		Sex             string `json:"sex"`
		Age             int    `json:"age"`
		MsgDetail       string `json:"msg_detail"`
		Status          string `json:"status"`
	} `json:"data,omitempty"`
}
type GroupEssenceMsg struct {
	Common
	Data []struct {
		SenderID     int64  `json:"sender_id"`
		SenderNick   string `json:"sender_nick"`
		SenderTime   int64  `json:"sender_time"`
		OperatorID   int64  `json:"operator_id"`
		OperatorNick string `json:"operator_nick"`
		OperatorTime int64  `json:"operator_time"`
		MessageID    int32  `json:"message_id"`
		MessageSeq   int32  `json:"message_seq"`
	} `json:"data,omitempty"`
}

type IsBlacklist struct {
	Common
	Data struct {
		Is bool `json:"is"`
	} `json:"data,omitempty"`
}

// ////////////////////////////////////////////////////////////消息

type CommonMsg struct {
	Common
	Data struct {
		MessageID int32 `json:"message_id"`
		Time      int64 `json:"time"`
	} `json:"data,omitempty"`
}

type PrivateMsg struct {
	Common
	Data struct {
		MessageID int32 `json:"message_id"`
		Time      int64 `json:"time"`
	} `json:"data,omitempty"`
}

type GroupMsg struct {
	Common
	Data struct {
		MessageID int32 `json:"message_id"`
		Time      int64 `json:"time"`
	} `json:"data,omitempty"`
}
type Sender struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	UID      string `json:"uid"`
}
type Msg struct {
	Common
	Data struct {
		Time        int64         `json:"time"`
		MessageType string        `json:"message_type"`
		MessageID   int32         `json:"message_id"`
		RealID      int32         `json:"real_id"`
		Sender      Sender        `json:"sender"`
		Message     []MessageItem `json:"message"`
		GroupID     int64         `json:"group_id"`
		TargetID    int64         `json:"target_id"`
		PeerID      int64         `json:"peer_id"`
	} `json:"data,omitempty"`
}

type HistoryMsg struct {
	Common
	Data struct {
		Messages []struct {
			Time        int64         `json:"time"`
			MessageType string        `json:"message_type"`
			MessageID   int32         `json:"message_id"`
			RealID      int32         `json:"real_id"`
			Sender      Sender        `json:"sender"`
			Message     []MessageItem `json:"message"`
			GroupID     int64         `json:"group_id"`
			TargetID    int64         `json:"target_id"`
			PeerID      int64         `json:"peer_id"`
		} `json:"messages"`
	} `json:"data,omitempty"`
}
type GroupHistoryMsg struct {
	Common
	Data struct {
		Messages []struct {
			Time        int64         `json:"time"`
			MessageType string        `json:"message_type"`
			MessageID   int32         `json:"message_id"`
			RealID      int32         `json:"real_id"`
			Sender      Sender        `json:"sender"`
			Message     []MessageItem `json:"message"`
			GroupID     int64         `json:"group_id"`
			TargetID    int64         `json:"target_id"`
			PeerID      int64         `json:"peer_id"`
		} `json:"messages"`
	} `json:"data,omitempty"`
}

type ForwardMsg struct {
	Common
	Data struct {
		Messages []struct {
			Time        int    `json:"time"`
			MessageType string `json:"message_type"`
			MessageId   int    `json:"message_id"`
			RealId      int    `json:"real_id"`
			Sender      struct {
				UserId   int    `json:"user_id"`
				Nickname string `json:"nickname"`
				Sex      string `json:"sex"`
				Age      int    `json:"age"`
				Uid      string `json:"uid"`
			} `json:"sender"`
			Message []struct {
				Type string `json:"type"`
				Data struct {
					Text string `json:"text"`
				} `json:"data,omitempty"`
			} `json:"message"`
			PeerId   int `json:"peer_id"`
			TargetId int `json:"target_id"`
		} `json:"messages"`
	} `json:"data,omitempty"`
}

// ////////////////////////////////////////////////////////////资源

type FileInfo struct {
	Common
	Data struct {
		Size     int64  `json:"size"`
		URL      string `json:"url"`
		Filename string `json:"filename"`
	} `json:"data,omitempty"`
}
type EnableSendImage struct {
	Common
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data,omitempty"`
}

type OcrImage struct {
	Common
	Data struct {
		Language      string `json:"language"`
		TextDetection struct {
			Text        string `json:"text"`
			Confidence  int32  `json:"confidence"`
			Coordinates []struct {
				X int32 `json:"x"`
				Y int32 `json:"y"`
			} `json:"coordinates"`
		} `json:"texts"`
	} `json:"data,omitempty"`
}
type AudioMsg struct {
	Common
	Data struct {
		File string `json:"file"`
		URL  string `json:"url"`
	} `json:"data,omitempty"`
}
type EnableSendAudio struct {
	Common
	Data struct {
		Yes bool `json:"yes"`
	} `json:"data,omitempty"`
}

// ////////////////////////////////////////////////////////////群聊

type GroupAvatar struct {
	Common
	Data struct {
		GroupID int64  `json:"group_id"`
		File    string `json:"file"`
		Cache   int    `json:"cache"`
	} `json:"data,omitempty"`
}
type GroupNotice struct {
	Common
	Data struct {
		SenderID    int64 `json:"sender_id"`
		PublishTime int64 `json:"publish_time"`
		Message     struct {
			Text   string `json:"text"`
			Images []struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				ID     string `json:"id"`
			} `json:"images"`
		} `json:"message"`
	} `json:"data,omitempty"`
}
type GroupProhibitedMember struct {
	Common
	Data []struct {
		UserId int64 `json:"user_id"`
		Time   int   `json:"time"`
	} `json:"data,omitempty"`
}

// ////////////////////////////////////////////////////////////文件

type SendPrivateFile struct {
	Common
	Data struct {
		MsgId  int    `json:"msg_id"`
		Bizid  int    `json:"bizid"`
		Md5    string `json:"md5"`
		Sha    string `json:"sha"`
		Sha3   string `json:"sha3"`
		FileId string `json:"file_id"`
	} `json:"data,omitempty"`
}
type SendGroupFile struct {
	Common
	Data struct {
		MsgId  int    `json:"msg_id"`
		Bizid  int    `json:"bizid"`
		Md5    string `json:"md5"`
		Sha    string `json:"sha"`
		Sha3   string `json:"sha3"`
		FileId string `json:"file_id"`
	} `json:"data,omitempty"`
}
type GroupFileSystemInfo struct {
	Common
	Data struct {
		FileCount  int32 `json:"file_count"`
		LimitCount int32 `json:"limit_count"`
		UsedSpace  int64 `json:"used_space"`
		TotalSpace int64 `json:"total_space"`
	} `json:"data,omitempty"`
}
type GroupRootFileList struct {
	Common
	Data struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	} `json:"data,omitempty"`
}
type File struct {
	GroupID       int32  `json:"group_id"`
	FileID        string `json:"file_id"`
	FileName      string `json:"file_name"`
	BusID         int32  `json:"busid"`
	FileSize      int64  `json:"file_size"`
	UploadTime    int64  `json:"upload_time"`
	DeadTime      int64  `json:"dead_time"`
	ModifyTime    int64  `json:"modify_time"`
	DownloadTimes int32  `json:"download_times"`
	Uploader      int64  `json:"uploader"`
	UploaderName  string `json:"uploader_name"`
	MD5           string `json:"md5"`
	SHA           string `json:"sha"`
	SHA3          string `json:"sha3"`
}
type Folder struct {
	GroupID        int32  `json:"group_id"`
	FolderID       string `json:"folder_id"`
	FolderName     string `json:"folder_name"`
	CreateTime     int64  `json:"create_time"`
	Creator        int64  `json:"creator"`
	CreatorName    string `json:"creator_name"`
	TotalFileCount int32  `json:"total_file_count"`
}
type GroupFiles struct {
	Common
	Data struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	} `json:"data,omitempty"`
}
type GroupFileUrl struct {
	Common
	Data struct {
		Url string `json:"url"`
	} `json:"data,omitempty"`
}
type UploadCache struct {
	Common
	Data struct {
		File string `json:"file"`
	} `json:"data,omitempty"`
}
type DownloadCache struct {
	Common
	Data struct {
		File string `json:"file"`
	} `json:"data,omitempty"`
}
type Battery struct {
	Common
	Data struct {
		Battery int   `json:"battery"`
		Scale   int64 `json:"scale"`
		Status  int   `json:"status"`
	} `json:"data,omitempty"`
}
type StartTime struct {
	Common
	Data int64 `json:"data,omitempty"`
}
type WeatherCityCode struct {
	Common
	Data []struct {
		Adcode   int    `json:"adcode"`
		Province string `json:"province"`
		City     string `json:"city"`
	} `json:"data,omitempty"`
}
type Weather struct {
	Common
	Data struct {
		WeekStore struct {
			WeatherInfo struct {
				AllAstro              []interface{} `json:"all_astro"`
				LifeindexForecastList []interface{} `json:"lifeindex_forecast_list"`
				WeeklyAstro           []interface{} `json:"weekly_astro"`
				Ret                   int           `json:"ret"`
				WeatherInfo           struct {
					Temper        string `json:"temper"`
					AirHumidity   string `json:"air_humidity"`
					WindPower     string `json:"wind_power"`
					WindDirect    string `json:"wind_direct"`
					Weather       string `json:"weather"`
					Pubtime       string `json:"pubtime"`
					Updatetime    int    `json:"updatetime"`
					WeatherType   string `json:"weather_type"`
					WeatherTypeId string `json:"weather_type_id"`
					TypeIdNew     int    `json:"type_id_new"`
					ConcreteType  int    `json:"concrete_type"`
					Sunrise       string `json:"sunrise"`
					Sunset        string `json:"sunset"`
				} `json:"weather_info"`
				AirInfo      interface{} `json:"air_info"`
				ForecastList struct {
					WeatherForecast []struct {
						DayWeather         string `json:"day_weather"`
						NightWeather       string `json:"night_weather"`
						DayTemper          string `json:"day_temper"`
						NightTemper        string `json:"night_temper"`
						DayWindDirect      string `json:"day_wind_direct"`
						NightWindDirect    string `json:"night_wind_direct"`
						DayWindPower       string `json:"day_wind_power"`
						NightWindPower     string `json:"night_wind_power"`
						SunriseTime        string `json:"sunrise_time"`
						SunsetTime         string `json:"sunset_time"`
						Pubtime            string `json:"pubtime"`
						Day                int    `json:"day"`
						DayWeatherType     string `json:"day_weather_type"`
						NightWeatherType   string `json:"night_weather_type"`
						DayWeatherTypeId   string `json:"day_weather_type_id"`
						NightWeatherTypeId string `json:"night_weather_type_id"`
						DayTypeIdNew       int    `json:"day_type_id_new"`
						DayConcreteType    int    `json:"day_concrete_type"`
						NightTypeIdNew     int    `json:"night_type_id_new"`
						NightConcreteType  int    `json:"night_concrete_type"`
						Pm                 string `json:"pm"`
						WindPowerDesc      string `json:"wind_power_desc"`
					} `json:"weatherForecast"`
					Updatetime     int    `json:"updatetime"`
					TomorrowPrompt string `json:"tomorrowPrompt"`
					WeeklyPrompt   string `json:"weeklyPrompt"`
				} `json:"forecast_list"`
				Forecast struct {
					DayWeather         string `json:"day_weather"`
					NightWeather       string `json:"night_weather"`
					DayTemper          string `json:"day_temper"`
					NightTemper        string `json:"night_temper"`
					DayWindDirect      string `json:"day_wind_direct"`
					NightWindDirect    string `json:"night_wind_direct"`
					DayWindPower       string `json:"day_wind_power"`
					NightWindPower     string `json:"night_wind_power"`
					SunriseTime        string `json:"sunrise_time"`
					SunsetTime         string `json:"sunset_time"`
					Pubtime            string `json:"pubtime"`
					Day                int    `json:"day"`
					DayWeatherType     string `json:"day_weather_type"`
					NightWeatherType   string `json:"night_weather_type"`
					DayWeatherTypeId   string `json:"day_weather_type_id"`
					NightWeatherTypeId string `json:"night_weather_type_id"`
					DayTypeIdNew       int    `json:"day_type_id_new"`
					DayConcreteType    int    `json:"day_concrete_type"`
					NightTypeIdNew     int    `json:"night_type_id_new"`
					NightConcreteType  int    `json:"night_concrete_type"`
					Pm                 string `json:"pm"`
					WindPowerDesc      string `json:"wind_power_desc"`
				} `json:"forecast"`
				HourinfoList interface{} `json:"hourinfo_list"`
				Almanac      string      `json:"almanac"`
				WarningList  struct {
					LstWarning   []interface{} `json:"lst_warning"`
					LastProcTime int           `json:"last_proc_time"`
				} `json:"warning_list"`
				Astro           interface{} `json:"astro"`
				City            string      `json:"city"`
				Area            string      `json:"area"`
				Adcode          int         `json:"adcode"`
				AreaId          int         `json:"area_id"`
				EnName          string      `json:"en_name"`
				UpdateTime      int         `json:"update_time"`
				TipsList        interface{} `json:"tips_list"`
				LifeindexList   interface{} `json:"lifeindex_list"`
				CurrentTime     int         `json:"current_time"`
				UserWeeklyAstro interface{} `json:"user_weekly_astro"`
				WeeklySummary   interface{} `json:"weekly_summary"`
			} `json:"weatherInfo"`
			Qrcode string `json:"qrcode"`
			Poster string `json:"poster"`
			Share  struct {
				Data struct {
					App    string `json:"app"`
					Config struct {
						Autosize int    `json:"autosize"`
						Ctime    int    `json:"ctime"`
						Forward  int    `json:"forward"`
						Round    int    `json:"round"`
						Token    string `json:"token"`
					} `json:"config"`
					Desc string `json:"desc"`
					Meta struct {
						Share struct {
							Adcode       int         `json:"adcode"`
							AirInfo      interface{} `json:"air_info"`
							Area         string      `json:"area"`
							City         string      `json:"city"`
							CurrentTime  int         `json:"current_time"`
							ForecastList struct {
								TomorrowPrompt  string `json:"tomorrowPrompt"`
								Updatetime      int    `json:"updatetime"`
								WeatherForecast []struct {
									Day                int    `json:"day"`
									DayConcreteType    int    `json:"day_concrete_type"`
									DayTemper          string `json:"day_temper"`
									DayTypeIdNew       int    `json:"day_type_id_new"`
									DayWeather         string `json:"day_weather"`
									DayWeatherType     string `json:"day_weather_type"`
									DayWeatherTypeId   string `json:"day_weather_type_id"`
									DayWindDirect      string `json:"day_wind_direct"`
									DayWindPower       string `json:"day_wind_power"`
									NightConcreteType  int    `json:"night_concrete_type"`
									NightTemper        string `json:"night_temper"`
									NightTypeIdNew     int    `json:"night_type_id_new"`
									NightWeather       string `json:"night_weather"`
									NightWeatherType   string `json:"night_weather_type"`
									NightWeatherTypeId string `json:"night_weather_type_id"`
									NightWindDirect    string `json:"night_wind_direct"`
									NightWindPower     string `json:"night_wind_power"`
									Pm                 string `json:"pm"`
									Pubtime            string `json:"pubtime"`
									SunriseTime        string `json:"sunrise_time"`
									SunsetTime         string `json:"sunset_time"`
									WindPowerDesc      string `json:"wind_power_desc"`
								} `json:"weatherForecast"`
								WeeklyPrompt string `json:"weeklyPrompt"`
							} `json:"forecast_list"`
							UpdateTime  int `json:"update_time"`
							WeatherInfo struct {
								AirHumidity   string `json:"air_humidity"`
								ConcreteType  int    `json:"concrete_type"`
								Pubtime       string `json:"pubtime"`
								Sunrise       string `json:"sunrise"`
								Sunset        string `json:"sunset"`
								Temper        string `json:"temper"`
								TypeIdNew     int    `json:"type_id_new"`
								Updatetime    int    `json:"updatetime"`
								Weather       string `json:"weather"`
								WeatherType   string `json:"weather_type"`
								WeatherTypeId string `json:"weather_type_id"`
								WindDirect    string `json:"wind_direct"`
								WindPower     string `json:"wind_power"`
							} `json:"weather_info"`
						} `json:"share"`
					} `json:"meta"`
					Prompt string `json:"prompt"`
					Ver    string `json:"ver"`
					View   string `json:"view"`
				} `json:"data,omitempty"`
				Code int `json:"code"`
			} `json:"share"`
		} `json:"weekStore"`
	} `json:"data,omitempty"`
}
