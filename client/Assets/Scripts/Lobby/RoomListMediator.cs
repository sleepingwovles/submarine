﻿using System.Collections.Generic;
using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class RoomListMediator : MediatorBase<RoomListView>, IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        CreateRoomCommand createRoomCommand;
        [Inject]
        GetRoomsCommand getRoomsCommand;
        [Inject]
        JoinIntoRoomCommand joinIntoRoomCommand;

        public void Initialize()
        {
            lobbyModel.Rooms.Subscribe(OnRoomsChange).AddTo(view);
            view.CreateRoomButtonClickedAsObservable().Subscribe(_ => OnCreateRoomButtonClick()).AddTo(view);
            view.UpdateRoomsButtonClickedAsObservable().Subscribe(_ => OnUpdateRoomsButtonClick()).AddTo(view);
        }

        void OnCreateRoomButtonClick()
        {
            createRoomCommand.Fire();
        }

        void OnUpdateRoomsButtonClick()
        {
            getRoomsCommand.Fire();
        }

        void OnRoomsChange(List<Type.Room> rooms)
        {
            view.ClearRooms();
            view.CreateRooms(rooms, OnRoomClick);
        }

        void OnRoomClick(Type.Room room)
        {
            joinIntoRoomCommand.Fire(room);
        }
    }
}
