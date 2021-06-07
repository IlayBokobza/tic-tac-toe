import { useDispatch } from 'react-redux'
import {setSocket} from '../redux/store'
import io from 'socket.io-client'
import { useState } from 'react'
import { Link, useHistory } from 'react-router-dom'
import Popup from './Popup'

const NavBar = () => {
    const [showPopup,setShowPopup] = useState(false)
    const dispatch = useDispatch()
    const history = useHistory()

    const joinGame = (id:string) => {
        //creates scoket ands stores it in redux
        const socket = io()
        dispatch(setSocket(socket))

        socket.emit('joinGame',id,(err:string) => {
            if(err){
                console.log(err)
                socket.disconnect()
                dispatch(setSocket(null))
                return
            }
            history.push('/gameroom')
        })
    }

    return (
        <div className="navbar">
            <h1>Tic Tac Toe</h1>
            <div className="links">
                <Link to="/">Home</Link>
                <button onClick={()=>{setShowPopup(true)}}>Join Game</button>
                <Link to="/create">Create Game</Link>
            </div>
            {showPopup && <Popup
                title="Join Game"
                text="Enter game id:"
                showInput={true}
                onlyAccept={false}
                acceptText="Join!"
                acceptFunc={joinGame}
                closeFunc={() => {setShowPopup(false)}}
            />}
        </div>
    )
}

export default NavBar