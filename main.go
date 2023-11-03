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
package main

import (
    "context"
    "encoding/base64"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "strings"

    "mywabot/system/message"

    _ "github.com/mattn/go-sqlite3"
    "github.com/mdp/qrterminal"
    "github.com/probandula/figlet4go"
    "github.com/joho/godotenv"

    "github.com/amiruldev20/waSocket"
    waProto "github.com/amiruldev20/waSocket/binary/proto"
    "github.com/amiruldev20/waSocket/store"
    "github.com/amiruldev20/waSocket/store/sqlstore"
    "github.com/amiruldev20/waSocket/types/events"
    "github.com/amiruldev20/waSocket/types"
    waLog "github.com/amiruldev20/waSocket/util/log"

    "google.golang.org/protobuf/proto"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error load file .env")
    }

    dbLog := waLog.Stdout("Database", "ERROR", true)
    dxz,
    err := base64.StdEncoding.DecodeString("TVlXQSBCT1Q=")
    if err != nil {
        panic("malformed input")
        log.Println(dxz)
    }
    container,
    err := sqlstore.New("sqlite3", "file:mywabot.db?_foreign_keys=on", dbLog)
    if err != nil {
        panic(err)
    }

    deviceStore,
    err := container.GetFirstDevice()
    if err != nil {
        panic(err)
    }
    clientLog := waLog.Stdout("Client", "ERROR", true)

    /* setting env */
        typeLogin := os.Getenv("TYPE_LOGIN")
    numberBot := os.Getenv("BOT_NUMBER")

    /* client */
        sock := waSocket.NewClient(deviceStore, clientLog)
    eventHandler := registerHandler(sock)
    sock.AddEventHandler(eventHandler)

    if sock.Store.ID == nil {
        if typeLogin == "code" {
            fmt.Println("You login with pairing code")
            fmt.Println("Bot Number: " + numberBot)

            err = sock.Connect()
            if err != nil {
                panic(err)
            }

            /* don't edit */
            code, err := sock.PairPhone(numberBot, true, waSocket.PairClientChrome, "Chrome (Linux)")

            if err != nil {
                fmt.Println(err)
                return
            }
            log.Println("Your Code: " + code)

        } else {
            qrChan,
            _ := sock.GetQRChannel(context.Background())

                err = sock.Connect()
            if err != nil {
                panic(err)
            }

            for evt := range qrChan {
                if evt.Event == "code" {
                    qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
                    dxz, err := base64.StdEncoding.DecodeString("TWFkZSBieSBBbWlydWwgRGV2LiBmb2xsb3cgSUcgQGFtaXJ1bC5kZXY=")

                    if err != nil {
                        panic("malformed input")
                    }
                    log.Println(string(dxz))

                    log.Println("Please scan this QR...")
                } else {
                    log.Println("Login successfully!!")
                }
            }
        }
    } else {

        /* Already logged in, just connect */
        err = sock.Connect()
        log.Println("Login Sucessfully!!")
        if err != nil {
            panic(err)
        }
    }

    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
<-c

    sock.Disconnect()
}

func init() {
    ascii := figlet4go.NewAsciiRender()
    dxz,
    err := base64.StdEncoding.DecodeString("TXl3YSBCT1QgYnkgd2FTb2NrZXQ=")
    if err != nil {
        panic("malformed input")
    }
    renderStr,
    _ := ascii.Render(string(dxz))
    store.DeviceProps.PlatformType = waProto.DeviceProps_FIREFOX.Enum()
    store.DeviceProps.Os = proto.String(string(dxz))
    fmt.Print(renderStr)
}

func registerHandler(sock * waSocket.Client) func(evt interface {}) {
    return func(evt interface {}) {
        switch v := evt.(type) {
            case *events.Message:
                if strings.HasPrefix(v.Info.ID, "BAE5") {
                    return
                }
                if v.Info.Chat.String() == "status@broadcast" {
                    sock.MarkRead([] types.MessageID {v.Info.ID}, v.Info.Timestamp, v.Info.Chat, v.Info.Sender)
                    fmt.Println("Berhasil melihat status", v.Info.PushName)
                }
                go message.Msg(sock, v)
                break
        }
    }
}