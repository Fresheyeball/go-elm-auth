module Main where

import Html exposing (..)
import Html.Events as Events
import SocketIO exposing (..)
import Task exposing (Task, andThen)
import Signal exposing (..)

messageKey : String
messageKey = "message"

socket : Task x Socket
socket =
  io "http://localhost:8000/socket" defaultOptions

onConnect : Mailbox Bool
onConnect = mailbox False

port connect : Task x ()
port connect =
  socket `andThen` connected onConnect.address

onInput : Mailbox String
onInput = mailbox ""

port send : Signal (Task x ())
port send =
  andThen socket << emit messageKey <~ onInput.signal

onMessage : Mailbox String
onMessage = mailbox ""

port recieve : Task x ()
port recieve =
  andThen socket <| on messageKey onMessage.address

view : Bool -> String -> Html
view status response = let
  status' = case status of
    True -> "Connected!"
    False -> "Not connected :("
  in div []
  [ h2 [] [ text status' ]
  , input [ Events.on "input" Events.targetValue (Signal.message onInput.address) ] []
  , br [] []
  , text ("factorial of the message length is " ++ response) ]

main : Signal Html
main = view <~ onConnect.signal ~ onMessage.signal
