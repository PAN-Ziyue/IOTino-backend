package MQTT

import (
    "IOTino/pkg"
    "fmt"
    "github.com/eclipse/paho.mqtt.golang/packets"
    "net"
    "time"
)

func MQTT() {
    listener, err := net.Listen("tcp", config.TCP_ADDR)
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            return // 终止程序
        }
        go BuildConnection(conn)
    }
}

func BuildConnection(conn net.Conn) {
    packet, err := packets.ReadPacket(conn)
    fmt.Println(packet)
    if err != nil {
        fmt.Println(err)
        return
    }
    if packet == nil {
        fmt.Println(err)
        return
    }
    msg, ok := packet.(*packets.ConnectPacket)
    if !ok {
        return
    }

    connack := packets.NewControlPacket(packets.Connack).(*packets.ConnackPacket)
    connack.SessionPresent = msg.CleanSession
    connack.ReturnCode = msg.Validate()

    if connack.ReturnCode != packets.Accepted {
        //err = connack.Write(conn)
        //if err != nil {
        //	return
        //}
        return
    }

    if err = connack.Write(conn); err != nil {
        //log.Error("send connack error, ", zap.Error(err), zap.String("clientID", msg.ClientIdentifier))
        return
    }

    ProcessMessage(conn)
}

func ProcessMessage(conn net.Conn) {

    timeout := time.Second * time.Duration(config.KEEP_ALIVE)

    for {
        if err := conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
            fmt.Println(err)
            return
        }

        packet, err := packets.ReadPacket(conn)
        if err != nil {
            fmt.Println(err)
            return
        }

        switch packet.(type) {
        case *packets.PublishPacket:
            ProcessPublish(conn, packet.(*packets.PublishPacket))
        case *packets.PubrelPacket:
            ProcessPubrel(conn, packet.(*packets.PubrelPacket))
        default:
            fmt.Println("Unsupported packets")
            fmt.Println(packet)
        }
    }
}

func ProcessPublish(conn net.Conn, packet *packets.PublishPacket) {
    fmt.Println(packet)
    switch packet.Qos {
    case 2:
        pubrec := packets.NewControlPacket(packets.Pubrec).(*packets.PubrecPacket)
        pubrec.MessageID = packet.MessageID
        if err := pubrec.Write(conn); err != nil {
            fmt.Println("[ERROR]")
            return
        }
    default:
        fmt.Println("Unsupported Qos")
    }
}

func ProcessPubrel(conn net.Conn, packet *packets.PubrelPacket) {
    pubcomp := packets.NewControlPacket(packets.Pubcomp).(*packets.PubcompPacket)
    pubcomp.MessageID = packet.MessageID
    if err := pubcomp.Write(conn); err != nil {
        fmt.Println("[ERROR]")
        return
    }
}
