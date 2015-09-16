module Main where

import Html exposing (..)
import SocketIO exposing (..)
import Task exposing (Task)
import Signal exposing (..)

socket : Task x Socket
socket =
  io "http://localhost:8000/socket" defaultOptions

onConnect : Mailbox Bool
onConnect = mailbox False

port connect : Task x ()
port connect = socket `Task.andThen` connected onConnect.address

main : Signal Html
main = let
  render status = case status of
    True -> "Connected!"
    False -> "Not connected :("
  in (\x -> h2 [] [x]) << text << render <~ onConnect.signal
