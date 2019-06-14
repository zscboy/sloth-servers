protoc --proto_path=../../proto --go_out=../../gosrc/lobbyserver/lobby  ../../proto/lobby.proto
protoc --proto_path=../../proto --go_out=../../gosrc/lobbyserver/lobby/club  ../../proto/lobby_club.proto
protoc --proto_path=../../proto --go_out=../../gosrc/lobbyserver/lobby/club  ../../proto/lobby_club_enum.proto
@pause