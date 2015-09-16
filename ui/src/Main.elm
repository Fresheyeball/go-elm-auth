module Main where

import Html exposing (..)
import Html.Events as Events
import SocketIO exposing (..)
import Task exposing (Task)
import Signal exposing (..)

socket : Task x Socket
socket =
  io "http://localhost:8000/socket" SocketIO.defaultOptions

onConnect : Mailbox Bool
onConnect = mailbox False

port connect : Task x ()
port connect = socket `Task.andThen` connected onConnect.address

onInput : Mailbox String
onInput = mailbox ""

view : Bool -> Html
view status = let
  status' = case status of
    True -> "Connected!"
    False -> "Not connected :("
  in div []
  [ h2 [] [ text status' ]
  , input [ Events.on "input" Events.targetValue (Signal.message onInput.address) ] [] ]

main : Signal Html
main = view <~ onConnect.signal
