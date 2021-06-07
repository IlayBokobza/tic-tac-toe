import { useDispatch } from 'react-redux'
import {setSocket} from '../redux/store'
import io from 'socket.io-client'
import { useEffect, useState } from 'react'
import { useHistory } from 'react-router'

const Create = () => {
    //creates scoket ands stores it in redux
    const dispatch = useDispatch()
    const socket = io()
    dispatch(setSocket(socket))
    const history = useHistory()

    //state
    const [id,setId]  = useState('')

    //aks the server to create game
    useEffect(() => {
        socket.emit('createGame',(id:string) => {
            setId(id)
        })

        socket.on('startGame',() => {
            history.push('/gameroom?start=true')
        })
    },[])

    return (
        <div className="create-page">
            <h1>New Game</h1>
            {id && (
                <div>
                    <h2>Game id: "{id}"</h2>
                    <p>Waiting other player...</p>  
                </div>
            )}
        </div>
    )
}

export default Create