package methods

import (
	"encoding/json"
	"gote/pkg/types"
	"gote/internal/utils"
	"gote/internal/utils/ctx"
)

const URL = "https://api.telegram.org/bot"


// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
//
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
//
// https://core.telegram.org/bots/api#getupdates
func GetUpdates (ctx ctx.CustomContext, param types.GetUpdates) ([]types.Update, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetUpdates"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.Update]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
//
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for webhooks: 443, 80, 88, 8443.If you&#39;re having any trouble setting up webhooks, please check out this amazing guide to webhooks.
//
// https://core.telegram.org/bots/api#setwebhook
func SetWebhook (ctx ctx.CustomContext, param types.SetWebhook) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetWebhook"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletewebhook
func DeleteWebhook (ctx ctx.CustomContext, param types.DeleteWebhook) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteWebhook"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
// 
// https://core.telegram.org/bots/api#getwebhookinfo
func GetWebhookInfo (ctx ctx.CustomContext, param types.GetWebhookInfo) (*types.WebhookInfo, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetWebhookInfo"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.WebhookInfo]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// A simple method for testing your bot&#39;s authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
// 
// https://core.telegram.org/bots/api#getme
func GetMe (ctx ctx.CustomContext, param types.GetMe) (*types.User, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMe"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.User]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
// 
// https://core.telegram.org/bots/api#logout
func LogOut (ctx ctx.CustomContext, param types.LogOut) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/LogOut"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn&#39;t launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
// 
// https://core.telegram.org/bots/api#close
func Close (ctx ctx.CustomContext, param types.Close) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/Close"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendmessage
func SendMessage (ctx ctx.CustomContext, param types.SendMessage) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to forward messages of any kind. Service messages and messages with protected content can&#39;t be forwarded. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#forwardmessage
func ForwardMessage (ctx ctx.CustomContext, param types.ForwardMessage) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/ForwardMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to forward multiple messages of any kind. If some of the specified messages can&#39;t be found or forwarded, they are skipped. Service messages and messages with protected content can&#39;t be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
// 
// https://core.telegram.org/bots/api#forwardmessages
func ForwardMessages (ctx ctx.CustomContext, param types.ForwardMessages) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/ForwardMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can&#39;t be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn&#39;t have a link to the original message. Returns the MessageId of the sent message on success.
// 
// https://core.telegram.org/bots/api#copymessage
func CopyMessage (ctx ctx.CustomContext, param types.CopyMessage) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/CopyMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to copy messages of any kind. If some of the specified messages can&#39;t be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can&#39;t be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don&#39;t have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
// 
// https://core.telegram.org/bots/api#copymessages
func CopyMessages (ctx ctx.CustomContext, param types.CopyMessages) (*types.MessageId, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/CopyMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.MessageId]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send photos. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendphoto
func SendPhoto (ctx ctx.CustomContext, param types.SendPhoto) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendPhoto"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendaudio
func SendAudio (ctx ctx.CustomContext, param types.SendAudio) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendAudio"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#senddocument
func SendDocument (ctx ctx.CustomContext, param types.SendDocument) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendDocument"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendvideo
func SendVideo (ctx ctx.CustomContext, param types.SendVideo) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendVideo"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendanimation
func SendAnimation (ctx ctx.CustomContext, param types.SendAnimation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendAnimation"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
// 
// https://core.telegram.org/bots/api#sendvoice
func SendVoice (ctx ctx.CustomContext, param types.SendVoice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendVoice"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendvideonote
func SendVideoNote (ctx ctx.CustomContext, param types.SendVideoNote) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendVideoNote"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send paid media. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendpaidmedia
func SendPaidMedia (ctx ctx.CustomContext, param types.SendPaidMedia) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendPaidMedia"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a group of photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Message objects that were sent is returned.
// 
// https://core.telegram.org/bots/api#sendmediagroup
func SendMediaGroup (ctx ctx.CustomContext, param types.SendMediaGroup) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendMediaGroup"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendlocation
func SendLocation (ctx ctx.CustomContext, param types.SendLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendLocation"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send information about a venue. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendvenue
func SendVenue (ctx ctx.CustomContext, param types.SendVenue) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendVenue"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send phone contacts. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendcontact
func SendContact (ctx ctx.CustomContext, param types.SendContact) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendContact"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a native poll. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendpoll
func SendPoll (ctx ctx.CustomContext, param types.SendPoll) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendPoll"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a checklist on behalf of a connected business account. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendchecklist
func SendChecklist (ctx ctx.CustomContext, param types.SendChecklist) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendChecklist"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#senddice
func SendDice (ctx ctx.CustomContext, param types.SendDice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendDice"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method when you need to tell the user that something is happening on the bot&#39;s side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// https://core.telegram.org/bots/api#sendchataction
func SendChatAction (ctx ctx.CustomContext, param types.SendChatAction) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SendChatAction"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the chosen reactions on a message. Service messages of some types can&#39;t be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can&#39;t use paid reactions. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmessagereaction
func SetMessageReaction (ctx ctx.CustomContext, param types.SetMessageReaction) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMessageReaction"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
// 
// https://core.telegram.org/bots/api#getuserprofilephotos
func GetUserProfilePhotos (ctx ctx.CustomContext, param types.GetUserProfilePhotos) (*types.UserProfilePhotos, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetUserProfilePhotos"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.UserProfilePhotos]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
// 
// https://core.telegram.org/bots/api#setuseremojistatus
func SetUserEmojiStatus (ctx ctx.CustomContext, param types.SetUserEmojiStatus) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetUserEmojiStatus"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot&lt;token&gt;/&lt;file_path&gt;, where &lt;file_path&gt; is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// 
// https://core.telegram.org/bots/api#getfile
func GetFile (ctx ctx.CustomContext, param types.GetFile) (*types.File, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetFile"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.File]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#banchatmember
func BanChatMember (ctx ctx.CustomContext, param types.BanChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/BanChatMember"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don&#39;t want this, use the parameter only_if_banned. Returns True on success.
// 
// https://core.telegram.org/bots/api#unbanchatmember
func UnbanChatMember (ctx ctx.CustomContext, param types.UnbanChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnbanChatMember"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
// 
// https://core.telegram.org/bots/api#restrictchatmember
func RestrictChatMember (ctx ctx.CustomContext, param types.RestrictChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/RestrictChatMember"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
// 
// https://core.telegram.org/bots/api#promotechatmember
func PromoteChatMember (ctx ctx.CustomContext, param types.PromoteChatMember) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/PromoteChatMember"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func SetChatAdministratorCustomTitle (ctx ctx.CustomContext, param types.SetChatAdministratorCustomTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatAdministratorCustomTitle"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won&#39;t be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#banchatsenderchat
func BanChatSenderChat (ctx ctx.CustomContext, param types.BanChatSenderChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/BanChatSenderChat"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#unbanchatsenderchat
func UnbanChatSenderChat (ctx ctx.CustomContext, param types.UnbanChatSenderChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnbanChatSenderChat"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatpermissions
func SetChatPermissions (ctx ctx.CustomContext, param types.SetChatPermissions) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatPermissions"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can&#39;t use invite links generated by other administrators. If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method. If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func ExportChatInviteLink (ctx ctx.CustomContext, param types.ExportChatInviteLink) (string, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	
	url := URL + ctx.Token + "/ExportChatInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return "", err
	}	
	
	var result utils.TGResponse[string]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#createchatinvitelink
func CreateChatInviteLink (ctx ctx.CustomContext, param types.CreateChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/CreateChatInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#editchatinvitelink
func EditChatInviteLink (ctx ctx.CustomContext, param types.EditChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditChatInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func CreateChatSubscriptionInviteLink (ctx ctx.CustomContext, param types.CreateChatSubscriptionInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/CreateChatSubscriptionInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func EditChatSubscriptionInviteLink (ctx ctx.CustomContext, param types.EditChatSubscriptionInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditChatSubscriptionInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
// 
// https://core.telegram.org/bots/api#revokechatinvitelink
func RevokeChatInviteLink (ctx ctx.CustomContext, param types.RevokeChatInviteLink) (*types.ChatInviteLink, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/RevokeChatInviteLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatInviteLink]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
// 
// https://core.telegram.org/bots/api#approvechatjoinrequest
func ApproveChatJoinRequest (ctx ctx.CustomContext, param types.ApproveChatJoinRequest) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ApproveChatJoinRequest"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
// 
// https://core.telegram.org/bots/api#declinechatjoinrequest
func DeclineChatJoinRequest (ctx ctx.CustomContext, param types.DeclineChatJoinRequest) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeclineChatJoinRequest"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a new profile photo for the chat. Photos can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatphoto
func SetChatPhoto (ctx ctx.CustomContext, param types.SetChatPhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatPhoto"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a chat photo. Photos can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletechatphoto
func DeleteChatPhoto (ctx ctx.CustomContext, param types.DeleteChatPhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteChatPhoto"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the title of a chat. Titles can&#39;t be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchattitle
func SetChatTitle (ctx ctx.CustomContext, param types.SetChatTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatTitle"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatdescription
func SetChatDescription (ctx ctx.CustomContext, param types.SetChatDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatDescription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to add a message to the list of pinned messages in a chat. In private chats and channel direct messages chats, all non-service messages can be pinned. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to pin messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#pinchatmessage
func PinChatMessage (ctx ctx.CustomContext, param types.PinChatMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/PinChatMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to remove a message from the list of pinned messages in a chat. In private chats and channel direct messages chats, all messages can be unpinned. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to unpin messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinchatmessage
func UnpinChatMessage (ctx ctx.CustomContext, param types.UnpinChatMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnpinChatMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a chat. In private chats and channel direct messages chats, no additional rights are required to unpin all pinned messages. Conversely, the bot must be an administrator with the &#39;can_pin_messages&#39; right or the &#39;can_edit_messages&#39; right to unpin all pinned messages in groups and channels respectively. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallchatmessages
func UnpinAllChatMessages (ctx ctx.CustomContext, param types.UnpinAllChatMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnpinAllChatMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
// 
// https://core.telegram.org/bots/api#leavechat
func LeaveChat (ctx ctx.CustomContext, param types.LeaveChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/LeaveChat"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
// 
// https://core.telegram.org/bots/api#getchat
func GetChat (ctx ctx.CustomContext, param types.GetChat) (*types.ChatFullInfo, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetChat"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatFullInfo]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a list of administrators in a chat, which aren&#39;t bots. Returns an Array of ChatMember objects.
// 
// https://core.telegram.org/bots/api#getchatadministrators
func GetChatAdministrators (ctx ctx.CustomContext, param types.GetChatAdministrators) ([]types.ChatMember, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetChatAdministrators"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.ChatMember]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
// 
// https://core.telegram.org/bots/api#getchatmembercount
func GetChatMemberCount (ctx ctx.CustomContext, param types.GetChatMemberCount) (int64, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return 0, err
	}
	
	url := URL + ctx.Token + "/GetChatMemberCount"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return 0, err
	}	
	
	var result utils.TGResponse[int64]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
// 
// https://core.telegram.org/bots/api#getchatmember
func GetChatMember (ctx ctx.CustomContext, param types.GetChatMember) (*types.ChatMember, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetChatMember"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatMember]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatstickerset
func SetChatStickerSet (ctx ctx.CustomContext, param types.SetChatStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatStickerSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletechatstickerset
func DeleteChatStickerSet (ctx ctx.CustomContext, param types.DeleteChatStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteChatStickerSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
// 
// https://core.telegram.org/bots/api#getforumtopiciconstickers
func GetForumTopicIconStickers (ctx ctx.CustomContext, param types.GetForumTopicIconStickers) ([]types.Sticker, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetForumTopicIconStickers"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.Sticker]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns information about the created topic as a ForumTopic object.
// 
// https://core.telegram.org/bots/api#createforumtopic
func CreateForumTopic (ctx ctx.CustomContext, param types.CreateForumTopic) (*types.ForumTopic, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/CreateForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ForumTopic]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit name and icon of a topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#editforumtopic
func EditForumTopic (ctx ctx.CustomContext, param types.EditForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/EditForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#closeforumtopic
func CloseForumTopic (ctx ctx.CustomContext, param types.CloseForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/CloseForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
// 
// https://core.telegram.org/bots/api#reopenforumtopic
func ReopenForumTopic (ctx ctx.CustomContext, param types.ReopenForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ReopenForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#deleteforumtopic
func DeleteForumTopic (ctx ctx.CustomContext, param types.DeleteForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func UnpinAllForumTopicMessages (ctx ctx.CustomContext, param types.UnpinAllForumTopicMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnpinAllForumTopicMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit the name of the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#editgeneralforumtopic
func EditGeneralForumTopic (ctx ctx.CustomContext, param types.EditGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/EditGeneralForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to close an open &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#closegeneralforumtopic
func CloseGeneralForumTopic (ctx ctx.CustomContext, param types.CloseGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/CloseGeneralForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to reopen a closed &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
// 
// https://core.telegram.org/bots/api#reopengeneralforumtopic
func ReopenGeneralForumTopic (ctx ctx.CustomContext, param types.ReopenGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ReopenGeneralForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to hide the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
// 
// https://core.telegram.org/bots/api#hidegeneralforumtopic
func HideGeneralForumTopic (ctx ctx.CustomContext, param types.HideGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/HideGeneralForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to unhide the &#39;General&#39; topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
// 
// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func UnhideGeneralForumTopic (ctx ctx.CustomContext, param types.UnhideGeneralForumTopic) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnhideGeneralForumTopic"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
// 
// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func UnpinAllGeneralForumTopicMessages (ctx ctx.CustomContext, param types.UnpinAllGeneralForumTopicMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UnpinAllGeneralForumTopicMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
//
// https://core.telegram.org/bots/api#answercallbackquery
func AnswerCallbackQuery (ctx ctx.CustomContext, param types.AnswerCallbackQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/AnswerCallbackQuery"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
// 
// https://core.telegram.org/bots/api#getuserchatboosts
func GetUserChatBoosts (ctx ctx.CustomContext, param types.GetUserChatBoosts) (*types.UserChatBoosts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetUserChatBoosts"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.UserChatBoosts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
// 
// https://core.telegram.org/bots/api#getbusinessconnection
func GetBusinessConnection (ctx ctx.CustomContext, param types.GetBusinessConnection) (*types.BusinessConnection, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetBusinessConnection"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.BusinessConnection]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the list of the bot&#39;s commands. See this manual for more details about bot commands. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmycommands
func SetMyCommands (ctx ctx.CustomContext, param types.SetMyCommands) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMyCommands"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete the list of the bot&#39;s commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemycommands
func DeleteMyCommands (ctx ctx.CustomContext, param types.DeleteMyCommands) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteMyCommands"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current list of the bot&#39;s commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren&#39;t set, an empty list is returned.
// 
// https://core.telegram.org/bots/api#getmycommands
func GetMyCommands (ctx ctx.CustomContext, param types.GetMyCommands) ([]types.BotCommand, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyCommands"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.BotCommand]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s name. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmyname
func SetMyName (ctx ctx.CustomContext, param types.SetMyName) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMyName"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
// 
// https://core.telegram.org/bots/api#getmyname
func GetMyName (ctx ctx.CustomContext, param types.GetMyName) (*types.BotName, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyName"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.BotName]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmydescription
func SetMyDescription (ctx ctx.CustomContext, param types.SetMyDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMyDescription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
// 
// https://core.telegram.org/bots/api#getmydescription
func GetMyDescription (ctx ctx.CustomContext, param types.GetMyDescription) (*types.BotDescription, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyDescription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.BotDescription]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s short description, which is shown on the bot&#39;s profile page and is sent together with the link when users share the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmyshortdescription
func SetMyShortDescription (ctx ctx.CustomContext, param types.SetMyShortDescription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMyShortDescription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
// 
// https://core.telegram.org/bots/api#getmyshortdescription
func GetMyShortDescription (ctx ctx.CustomContext, param types.GetMyShortDescription) (*types.BotShortDescription, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyShortDescription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.BotShortDescription]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the bot&#39;s menu button in a private chat, or the default menu button. Returns True on success.
// 
// https://core.telegram.org/bots/api#setchatmenubutton
func SetChatMenuButton (ctx ctx.CustomContext, param types.SetChatMenuButton) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetChatMenuButton"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current value of the bot&#39;s menu button in a private chat, or the default menu button. Returns MenuButton on success.
// 
// https://core.telegram.org/bots/api#getchatmenubutton
func GetChatMenuButton (ctx ctx.CustomContext, param types.GetChatMenuButton) (*types.MenuButton, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetChatMenuButton"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.MenuButton]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the default administrator rights requested by the bot when it&#39;s added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func SetMyDefaultAdministratorRights (ctx ctx.CustomContext, param types.SetMyDefaultAdministratorRights) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetMyDefaultAdministratorRights"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
// 
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func GetMyDefaultAdministratorRights (ctx ctx.CustomContext, param types.GetMyDefaultAdministratorRights) (*types.ChatAdministratorRights, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyDefaultAdministratorRights"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.ChatAdministratorRights]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a Gifts object.
// 
// https://core.telegram.org/bots/api#getavailablegifts
func GetAvailableGifts (ctx ctx.CustomContext, param types.GetAvailableGifts) (*types.Gifts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetAvailableGifts"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Gifts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Sends a gift to the given user or channel chat. The gift can&#39;t be converted to Telegram Stars by the receiver. Returns True on success.
// 
// https://core.telegram.org/bots/api#sendgift
func SendGift (ctx ctx.CustomContext, param types.SendGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SendGift"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Gifts a Telegram Premium subscription to the given user. Returns True on success.
// 
// https://core.telegram.org/bots/api#giftpremiumsubscription
func GiftPremiumSubscription (ctx ctx.CustomContext, param types.GiftPremiumSubscription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/GiftPremiumSubscription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#verifyuser
func VerifyUser (ctx ctx.CustomContext, param types.VerifyUser) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/VerifyUser"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#verifychat
func VerifyChat (ctx ctx.CustomContext, param types.VerifyChat) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/VerifyChat"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#removeuserverification
func RemoveUserVerification (ctx ctx.CustomContext, param types.RemoveUserVerification) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/RemoveUserVerification"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#removechatverification
func RemoveChatVerification (ctx ctx.CustomContext, param types.RemoveChatVerification) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/RemoveChatVerification"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#readbusinessmessage
func ReadBusinessMessage (ctx ctx.CustomContext, param types.ReadBusinessMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ReadBusinessMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletebusinessmessages
func DeleteBusinessMessages (ctx ctx.CustomContext, param types.DeleteBusinessMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteBusinessMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountname
func SetBusinessAccountName (ctx ctx.CustomContext, param types.SetBusinessAccountName) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetBusinessAccountName"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountusername
func SetBusinessAccountUsername (ctx ctx.CustomContext, param types.SetBusinessAccountUsername) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetBusinessAccountUsername"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountbio
func SetBusinessAccountBio (ctx ctx.CustomContext, param types.SetBusinessAccountBio) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetBusinessAccountBio"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func SetBusinessAccountProfilePhoto (ctx ctx.CustomContext, param types.SetBusinessAccountProfilePhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetBusinessAccountProfilePhoto"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func RemoveBusinessAccountProfilePhoto (ctx ctx.CustomContext, param types.RemoveBusinessAccountProfilePhoto) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/RemoveBusinessAccountProfilePhoto"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func SetBusinessAccountGiftSettings (ctx ctx.CustomContext, param types.SetBusinessAccountGiftSettings) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetBusinessAccountGiftSettings"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns StarAmount on success.
// 
// https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func GetBusinessAccountStarBalance (ctx ctx.CustomContext, param types.GetBusinessAccountStarBalance) (*types.StarAmount, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetBusinessAccountStarBalance"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.StarAmount]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Transfers Telegram Stars from the business account balance to the bot&#39;s balance. Requires the can_transfer_stars business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#transferbusinessaccountstars
func TransferBusinessAccountStars (ctx ctx.CustomContext, param types.TransferBusinessAccountStars) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/TransferBusinessAccountStars"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns OwnedGifts on success.
// 
// https://core.telegram.org/bots/api#getbusinessaccountgifts
func GetBusinessAccountGifts (ctx ctx.CustomContext, param types.GetBusinessAccountGifts) (*types.OwnedGifts, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetBusinessAccountGifts"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.OwnedGifts]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Converts a given regular gift to Telegram Stars. Requires the can_convert_gifts_to_stars business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#convertgifttostars
func ConvertGiftToStars (ctx ctx.CustomContext, param types.ConvertGiftToStars) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ConvertGiftToStars"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Upgrades a given regular gift to a unique gift. Requires the can_transfer_and_upgrade_gifts business bot right. Additionally requires the can_transfer_stars business bot right if the upgrade is paid. Returns True on success.
// 
// https://core.telegram.org/bots/api#upgradegift
func UpgradeGift (ctx ctx.CustomContext, param types.UpgradeGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/UpgradeGift"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Transfers an owned unique gift to another user. Requires the can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars business bot right if the transfer is paid. Returns True on success.
// 
// https://core.telegram.org/bots/api#transfergift
func TransferGift (ctx ctx.CustomContext, param types.TransferGift) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/TransferGift"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
// 
// https://core.telegram.org/bots/api#poststory
func PostStory (ctx ctx.CustomContext, param types.PostStory) (*types.Story, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/PostStory"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Story]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
// 
// https://core.telegram.org/bots/api#editstory
func EditStory (ctx ctx.CustomContext, param types.EditStory) (*types.Story, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditStory"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Story]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestory
func DeleteStory (ctx ctx.CustomContext, param types.DeleteStory) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteStory"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagetext
func EditMessageText (ctx ctx.CustomContext, param types.EditMessageText) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageText"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagecaption
func EditMessageCaption (ctx ctx.CustomContext, param types.EditMessageCaption) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageCaption"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo or a video otherwise. When an inline message is edited, a new file can&#39;t be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagemedia
func EditMessageMedia (ctx ctx.CustomContext, param types.EditMessageMedia) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageMedia"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// 
// https://core.telegram.org/bots/api#editmessagelivelocation
func EditMessageLiveLocation (ctx ctx.CustomContext, param types.EditMessageLiveLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageLiveLocation"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
// 
// https://core.telegram.org/bots/api#stopmessagelivelocation
func StopMessageLiveLocation (ctx ctx.CustomContext, param types.StopMessageLiveLocation) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/StopMessageLiveLocation"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit a checklist on behalf of a connected business account. On success, the edited Message is returned.
// 
// https://core.telegram.org/bots/api#editmessagechecklist
func EditMessageChecklist (ctx ctx.CustomContext, param types.EditMessageChecklist) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageChecklist"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
// 
// https://core.telegram.org/bots/api#editmessagereplymarkup
func EditMessageReplyMarkup (ctx ctx.CustomContext, param types.EditMessageReplyMarkup) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/EditMessageReplyMarkup"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
// 
// https://core.telegram.org/bots/api#stoppoll
func StopPoll (ctx ctx.CustomContext, param types.StopPoll) (*types.Poll, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/StopPoll"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Poll]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to approve a suggested post in a direct messages chat. The bot must have the &#39;can_post_messages&#39; administrator right in the corresponding channel chat. Returns True on success.
// 
// https://core.telegram.org/bots/api#approvesuggestedpost
func ApproveSuggestedPost (ctx ctx.CustomContext, param types.ApproveSuggestedPost) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ApproveSuggestedPost"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to decline a suggested post in a direct messages chat. The bot must have the &#39;can_manage_direct_messages&#39; administrator right in the corresponding channel chat. Returns True on success.
// 
// https://core.telegram.org/bots/api#declinesuggestedpost
func DeclineSuggestedPost (ctx ctx.CustomContext, param types.DeclineSuggestedPost) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeclineSuggestedPost"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Service messages about a supergroup, channel, or forum topic creation can&#39;t be deleted.- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages administrator right in a supergroup or a channel, it can delete any message there.- If the bot has can_manage_direct_messages administrator right in a channel, it can delete any message in the corresponding direct messages chat.Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemessage
func DeleteMessage (ctx ctx.CustomContext, param types.DeleteMessage) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete multiple messages simultaneously. If some of the specified messages can&#39;t be found, they are skipped. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletemessages
func DeleteMessages (ctx ctx.CustomContext, param types.DeleteMessages) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteMessages"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendsticker
func SendSticker (ctx ctx.CustomContext, param types.SendSticker) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendSticker"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
// 
// https://core.telegram.org/bots/api#getstickerset
func GetStickerSet (ctx ctx.CustomContext, param types.GetStickerSet) (*types.StickerSet, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetStickerSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.StickerSet]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
// 
// https://core.telegram.org/bots/api#getcustomemojistickers
func GetCustomEmojiStickers (ctx ctx.CustomContext, param types.GetCustomEmojiStickers) ([]types.Sticker, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetCustomEmojiStickers"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.Sticker]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
// 
// https://core.telegram.org/bots/api#uploadstickerfile
func UploadStickerFile (ctx ctx.CustomContext, param types.UploadStickerFile) (*types.File, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/UploadStickerFile"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.File]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
// 
// https://core.telegram.org/bots/api#createnewstickerset
func CreateNewStickerSet (ctx ctx.CustomContext, param types.CreateNewStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/CreateNewStickerSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
// 
// https://core.telegram.org/bots/api#addstickertoset
func AddStickerToSet (ctx ctx.CustomContext, param types.AddStickerToSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/AddStickerToSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickerpositioninset
func SetStickerPositionInSet (ctx ctx.CustomContext, param types.SetStickerPositionInSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerPositionInSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestickerfromset
func DeleteStickerFromSet (ctx ctx.CustomContext, param types.DeleteStickerFromSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteStickerFromSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
// 
// https://core.telegram.org/bots/api#replacestickerinset
func ReplaceStickerInSet (ctx ctx.CustomContext, param types.ReplaceStickerInSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/ReplaceStickerInSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickeremojilist
func SetStickerEmojiList (ctx ctx.CustomContext, param types.SetStickerEmojiList) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerEmojiList"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickerkeywords
func SetStickerKeywords (ctx ctx.CustomContext, param types.SetStickerKeywords) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerKeywords"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickermaskposition
func SetStickerMaskPosition (ctx ctx.CustomContext, param types.SetStickerMaskPosition) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerMaskPosition"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the title of a created sticker set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickersettitle
func SetStickerSetTitle (ctx ctx.CustomContext, param types.SetStickerSetTitle) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerSetTitle"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setstickersetthumbnail
func SetStickerSetThumbnail (ctx ctx.CustomContext, param types.SetStickerSetThumbnail) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetStickerSetThumbnail"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
// 
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func SetCustomEmojiStickerSetThumbnail (ctx ctx.CustomContext, param types.SetCustomEmojiStickerSetThumbnail) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetCustomEmojiStickerSetThumbnail"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
// 
// https://core.telegram.org/bots/api#deletestickerset
func DeleteStickerSet (ctx ctx.CustomContext, param types.DeleteStickerSet) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/DeleteStickerSet"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
// 
// https://core.telegram.org/bots/api#answerinlinequery
func AnswerInlineQuery (ctx ctx.CustomContext, param types.AnswerInlineQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/AnswerInlineQuery"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
// 
// https://core.telegram.org/bots/api#answerwebappquery
func AnswerWebAppQuery (ctx ctx.CustomContext, param types.AnswerWebAppQuery) (*types.SentWebAppMessage, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/AnswerWebAppQuery"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.SentWebAppMessage]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
// 
// https://core.telegram.org/bots/api#savepreparedinlinemessage
func SavePreparedInlineMessage (ctx ctx.CustomContext, param types.SavePreparedInlineMessage) (*types.PreparedInlineMessage, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SavePreparedInlineMessage"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.PreparedInlineMessage]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send invoices. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendinvoice
func SendInvoice (ctx ctx.CustomContext, param types.SendInvoice) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendInvoice"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
// 
// https://core.telegram.org/bots/api#createinvoicelink
func CreateInvoiceLink (ctx ctx.CustomContext, param types.CreateInvoiceLink) (string, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	
	url := URL + ctx.Token + "/CreateInvoiceLink"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return "", err
	}	
	
	var result utils.TGResponse[string]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
// 
// https://core.telegram.org/bots/api#answershippingquery
func AnswerShippingQuery (ctx ctx.CustomContext, param types.AnswerShippingQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/AnswerShippingQuery"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
// 
// https://core.telegram.org/bots/api#answerprecheckoutquery
func AnswerPreCheckoutQuery (ctx ctx.CustomContext, param types.AnswerPreCheckoutQuery) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/AnswerPreCheckoutQuery"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// A method to get the current Telegram Stars balance of the bot. Requires no parameters. On success, returns a StarAmount object.
// 
// https://core.telegram.org/bots/api#getmystarbalance
func GetMyStarBalance (ctx ctx.CustomContext, param types.GetMyStarBalance) (*types.StarAmount, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetMyStarBalance"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.StarAmount]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Returns the bot&#39;s Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
// 
// https://core.telegram.org/bots/api#getstartransactions
func GetStarTransactions (ctx ctx.CustomContext, param types.GetStarTransactions) (*types.StarTransactions, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetStarTransactions"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.StarTransactions]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
// 
// https://core.telegram.org/bots/api#refundstarpayment
func RefundStarPayment (ctx ctx.CustomContext, param types.RefundStarPayment) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/RefundStarPayment"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
// 
// https://core.telegram.org/bots/api#edituserstarsubscription
func EditUserStarSubscription (ctx ctx.CustomContext, param types.EditUserStarSubscription) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/EditUserStarSubscription"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// 
// https://core.telegram.org/bots/api#setpassportdataerrors
func SetPassportDataErrors (ctx ctx.CustomContext, param types.SetPassportDataErrors) (bool, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return false, err
	}
	
	url := URL + ctx.Token + "/SetPassportDataErrors"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return false, err
	}	
	
	var result utils.TGResponse[bool]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to send a game. On success, the sent Message is returned.
// 
// https://core.telegram.org/bots/api#sendgame
func SendGame (ctx ctx.CustomContext, param types.SendGame) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SendGame"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user&#39;s current score in the chat and force is False.
// 
// https://core.telegram.org/bots/api#setgamescore
func SetGameScore (ctx ctx.CustomContext, param types.SetGameScore) (*types.Message, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/SetGameScore"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[*types.Message]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
//
// https://core.telegram.org/bots/api#getgamehighscores
func GetGameHighScores (ctx ctx.CustomContext, param types.GetGameHighScores) ([]types.GameHighScore, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	
	url := URL + ctx.Token + "/GetGameHighScores"
	resp, err := utils.RequestWithContext(ctx.GoContext, url, data)
	if err != nil {
		return nil, err
	}	
	
	var result utils.TGResponse[[]types.GameHighScore]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return result.Result, err
	}	

	return result.Result, nil
}

