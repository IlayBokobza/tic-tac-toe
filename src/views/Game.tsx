import { CSSProperties, useEffect, useState } from "react"
import {socket} from '../service/socket'
import { useHistory } from "react-router"
import qs from 'qs'

const Game = () => {
    //gets query 
    const {start} = qs.parse(window.location.href.split('?')[1])
    const history = useHistory()
    //vars
    const [board,setBoard] = useState([
        [0,0,0],
        [0,0,0],
        [0,0,0],
    ])

    const [isMyTurn,setIsMyTurn] = useState((start && start === 'true') ? true:false)

    useEffect(() => {
        //socket events
        socket.on('madeTurn',(cords:number[]) => {
            const [y,x] = cords
            const tempBoard = [...board]
            tempBoard[y][x] = (start) ? 2:1
            setBoard(tempBoard)
            setIsMyTurn(true)
        })

        socket.on('win',(winPosJson:string) => {
            setIsMyTurn(false)
            const winPoses:number[][] = JSON.parse(winPosJson)
    
            winPoses.forEach(winPos => {
                const [x,y] = winPos
                const col = document.getElementById(`${y}|${x}`)!
                
                col.style.backgroundColor = '#25d40b'
            })
        })

        socket.on('gameOver',() => {
            history.push('/')
        })

        //on unload
        return () => {
            socket.emit('endGame')
        }
    },[])

    //checks if user is in game
    socket.emit('isInGame',(err:string) => {
        if(err){
            console.log(err)
            history.push('/')
            return (<div></div>)
        }
    })

    //gets correct jsx for spot in bord
    const getBoardElement = (boardValue:number) => {
        if(boardValue === 1){
            return <span className="icon material-icons">close</span>
        }else if(boardValue === 2){
            return <span className="icon material-icons">radio_button_unchecked</span>
        }else{
            return ""
        }
    }

    const handleClick = ({y,x}:{y:number,x:number}) => {
        //if its not his trun stop func
        if(!isMyTurn) return

        //if spot is taken stop func
        if(board[y][x] > 0) return


        socket.emit('madeTurn',[y,x],(err:string) => {
            if(err){
                console.warn(err)
                return
            }
            const tempBoard = [...board]
            tempBoard[y][x] = (start) ? 1:2
            setBoard(tempBoard)
            setIsMyTurn(false)
        })
    }

    const checkBoredPosForStyle = ({x,y}:{x:number,y:number}):CSSProperties => {
        if(board[y][x] > 0){
            return {
                cursor:'default'
            }
        }else {
            return {
                cursor:'pointer'
            }
        }
    }

    return (
        <div id="game-page" className="game-page">
            <h1>Game Room</h1>
            <div className="game">
                {board.map((row,y) => (
                    <div className="row" key={y}>
                        {row.map((col,x) => (
                            <div key={x} id={`${y}|${x}`} className="col" style={checkBoredPosForStyle({x,y})} onClick={() => {handleClick({y,x})}}>{getBoardElement(col)}</div>
                        ))}
                    </div>
                ))}
            </div>
        </div>
    )
}

export default Game