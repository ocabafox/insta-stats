package models

// IGJsonData ...
type IGJsonData struct {
	ActivityCounts interface{} `json:"activity_counts"`
	Config         struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
	} `json:"config"`
	SupportsEs6  bool   `json:"supports_es6"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string `json:"logging_page_id"`
			ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
			Graphql               struct {
				User struct {
					Biography              string      `json:"biography"`
					BlockedByViewer        bool        `json:"blocked_by_viewer"`
					CountryBlock           bool        `json:"country_block"`
					ExternalURL            interface{} `json:"external_url"`
					ExternalURLLinkshimmed interface{} `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					FollowedByViewer bool `json:"followed_by_viewer"`
					EdgeFollow       struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
					FollowsViewer            bool        `json:"follows_viewer"`
					FullName                 string      `json:"full_name"`
					HasBlockedViewer         bool        `json:"has_blocked_viewer"`
					HasRequestedViewer       bool        `json:"has_requested_viewer"`
					ID                       string      `json:"id"`
					IsPrivate                bool        `json:"is_private"`
					IsVerified               bool        `json:"is_verified"`
					MutualFollowers          interface{} `json:"mutual_followers"`
					ProfilePicURL            string      `json:"profile_pic_url"`
					ProfilePicURLHd          string      `json:"profile_pic_url_hd"`
					RequestedByViewer        bool        `json:"requested_by_viewer"`
					Username                 string      `json:"username"`
					ConnectedFbPage          interface{} `json:"connected_fb_page"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename           string `json:"__typename"`
								ID                 string `json:"id"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								Shortcode          string `json:"shortcode"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								Dimensions       struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								GatingInfo   interface{} `json:"gating_info"`
								MediaPreview string      `json:"media_preview"`
								Owner        struct {
									ID string `json:"id"`
								} `json:"owner"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								IsVideo bool `json:"is_video"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
					EdgeSavedMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_saved_media"`
					EdgeMediaCollections struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_collections"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Gatekeepers struct {
		Ld    bool `json:"ld"`
		Seo   bool `json:"seo"`
		Seoht bool `json:"seoht"`
	} `json:"gatekeepers"`
	Knobs struct {
		AcctNtb int `json:"acct:ntb"`
	} `json:"knobs"`
	Qe struct {
		DashForVod struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"dash_for_vod"`
		Aysf struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"aysf"`
		Bc3L struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bc3l"`
		CommentReporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"comment_reporting"`
		DirectReporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"direct_reporting"`
		Reporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reporting"`
		MediaReporting struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"media_reporting"`
		AccRecoveryLink struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"acc_recovery_link"`
		Notif struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"notif"`
		DrctNav struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"drct_nav"`
		PlPivotLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"pl_pivot_li"`
		PlPivotLo struct {
			G string `json:"g"`
			P struct {
				ShowPivot string `json:"show_pivot"`
			} `json:"p"`
		} `json:"pl_pivot_lo"`
		Four04AsReact struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"404_as_react"`
		AccRecovery struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"acc_recovery"`
		ClientGql struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"client_gql"`
		Collections struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"collections"`
		CommentTa struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"comment_ta"`
		Connections struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"connections"`
		DiscPpl struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"disc_ppl"`
		Embeds struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"embeds"`
		EbdsimLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebdsim_li"`
		EbdsimLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"ebdsim_lo"`
		Es6 struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"es6"`
		ExitStoryCreation struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"exit_story_creation"`
		Fs struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"fs"`
		GdprLoggedOut struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"gdpr_logged_out"`
		Appsell struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"appsell"`
		Imgopt struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"imgopt"`
		FollowButton struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"follow_button"`
		Loggedout struct {
			G string `json:"g"`
			P struct {
				NewCta             string `json:"new_cta"`
				RemoveUpsellBanner string `json:"remove_upsell_banner"`
				UpdateNav          string `json:"update_nav"`
			} `json:"p"`
		} `json:"loggedout"`
		LoggedoutUpsell struct {
			G string `json:"g"`
			P struct {
				HasNewLoggedoutUpsellContent string `json:"has_new_loggedout_upsell_content"`
			} `json:"p"`
		} `json:"loggedout_upsell"`
		UsLi struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"us_li"`
		Msisdn struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"msisdn"`
		BgSync struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"bg_sync"`
		Onetaplogin struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"onetaplogin"`
		OnetaploginUserbased struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"onetaplogin_userbased"`
		LoginPoe struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"login_poe"`
		PrvcyTggl struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"prvcy_tggl"`
		PrivateLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"private_lo"`
		ProfilePhotoNuxFbcV2 struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"profile_photo_nux_fbc_v2"`
		PushNotifications struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"push_notifications"`
		Reg struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reg"`
		RegVp struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"reg_vp"`
		FeedVp struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"feed_vp"`
		ReportHaf struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"report_haf"`
		ReportMedia struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"report_media"`
		ReportProfile struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"report_profile"`
		Save struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"save"`
		Sidecar struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"sidecar"`
		SuUniverse struct {
			G string `json:"g"`
			P struct {
				UseAutocompleteLogin string `json:"use_autocomplete_login"`
			} `json:"p"`
		} `json:"su_universe"`
		Stale struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stale"`
		StoriesLo struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stories_lo"`
		Stories struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"stories"`
		TpPblshr struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"tp_pblshr"`
		Video struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"video"`
		GdprSettings struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"gdpr_settings"`
		GdprEuTos struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"gdpr_eu_tos"`
		GdprRowTos struct {
			G string `json:"g"`
			P struct {
			} `json:"p"`
		} `json:"gdpr_row_tos"`
	} `json:"qe"`
	Hostname                     string `json:"hostname"`
	DisplayPropertiesServerGuess struct {
		PixelRatio     float64 `json:"pixel_ratio"`
		ViewportWidth  int     `json:"viewport_width"`
		ViewportHeight int     `json:"viewport_height"`
		Orientation    string  `json:"orientation"`
	} `json:"display_properties_server_guess"`
	EnvironmentSwitcherVisibleServerGuess bool   `json:"environment_switcher_visible_server_guess"`
	Platform                              string `json:"platform"`
	RhxGis                                string `json:"rhx_gis"`
	Nonce                                 string `json:"nonce"`
	IsBot                                 bool   `json:"is_bot"`
	ZeroData                              struct {
	} `json:"zero_data"`
	RolloutHash    string `json:"rollout_hash"`
	BundleVariant  string `json:"bundle_variant"`
	ProbablyHasApp bool   `json:"probably_has_app"`
	ShowAppInstall bool   `json:"show_app_install"`
}
