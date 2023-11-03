/*
###################################
# Name: Mywa BOT                  #
# Version: 1.0.1                  #
# Developer: Amirul Dev           #
# Library: waSocket               #
# Contact: 085157489446           #
###################################
# Thanks to: 
# Vnia
*/
package message

import (
    "fmt"
    "os"
    "strings"

    "github.com/joho/godotenv"
    
    "mywabot/system/lib"

    // "google.golang.org/protobuf/proto"

    "github.com/amiruldev20/waSocket"
    //waProto "github.com/amiruldev20/waSocket/binary/proto"
    "github.com/amiruldev20/waSocket/types/events"
)

var (
    prefix = "."
    self = false 
	owner = "6285157489446"
)

func Msg(sock * waSocket.Client, msg * events.Message) {

    err := godotenv.Load()
    if err != nil {
        panic("Error load file .env")
    }
    //botNumber := os.Getenv("BOT_NUMBER")

    /* my function */
    m := lib.NewSimp(sock, msg)
    //from := msg.Info.Chat
    sender := msg.Info.Sender.String()
    pushName := msg.Info.PushName
    isOwner := strings.Contains(sender, owner)
    //isAdmin := m.GetGroupAdmin(from, sender)
    //isBotAdm := m.GetGroupAdmin(from, botNumber + "@s.whatsapp.net")
    //isGroup := msg.Info.IsGroup
    args := strings.Split(m.GetCMD(), " ")
    command := strings.ToLower(args[0])
    //query := strings.Join(args[1: ], ` `)
    //extended := msg.Message.GetExtendedTextMessage()
    //quotedMsg := extended.GetContextInfo().GetQuotedMessage()
    //quotedImage := quotedMsg.GetImageMessage()
    //quotedVideo := quotedMsg.GetVideoMessage()
    //quotedSticker := quotedMsg.GetStickerMessage()
    
	// Self
    if self && !isOwner {
        return
    }

    //-- CONSOLE LOG
    fmt.Println("\n===============================\nNAME: " + pushName + "\nJID: " + sender + "\nTYPE: " + msg.Info.Type + "\nMessage: " + m.GetCMD() + "")
    fmt.Println(m.Msg.Message.GetPollUpdateMessage().GetMetadata())
    switch command {

        /* panggil bot */
        case "bot":
            m.Reply(`Bot aktif *` + pushName + `*`)
            m.React("🤖")

       
		// end
		}
}