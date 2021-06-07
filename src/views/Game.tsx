import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { useHistory } from "react-router"
import qs from 'qs'

const sokcetEvents = (socket:any,board:number[][],setBoard:Function) => {
        //socket events
        socket.on('madeTurn',(cords:number[]) => {
            console.log('here')
            const [y,x] = cords
            const tempBoard = [...board]
            tempBoard[y][x] = 2
            setBoard(tempBoard)
        })
}

const Game = () => {
    //gets query 
    const {start} = qs.parse(window.location.href.split('?')[1])

    const socket = useSelector((state:any) => state?.storeReducer?.socket?.payload)
    const history = useHistory()
    //vars
    const [board,setBoard] = useState([
        [0,0,0],
        [0,0,0],
        [0,0,0],
    ])

    const [isMyTurn,setIsMyTurn] = useState((start && start === 'true') ? true:false)


    //code

    if(!socket){
        history.push('/')
        return <div></div>
    }

    sokcetEvents(socket,board,setBoard)

    //gets correct jsx for spot in bord
    const getBoardElement = (boardValue:number) => {
        if(boardValue === 1){
            return "X"
        }else if(boardValue === 2){
            return "0"
        }else{
            return ""
        }
    }

    const handleClick = (e:any) => {
        console.log(isMyTurn)
        if(!isMyTurn){return}

        const id = e.target.id as string
        let [yStr,xStr] = id.split('|')
        const y = parseInt(yStr),x = parseInt(xStr)
        setIsMyTurn(false)
        
        const tempBoard = [...board]
        tempBoard[y][x] = 1
        setBoard(tempBoard)
        socket.emit('madeTurn',[y,x],(err:string) => {
            console.warn(err)
        })
    }

    return (
        <div className="game-page">
            <h1>Game Room</h1>
            <div className="game">
                {board.map((row,y) => (
                    <div className="row" key={y}>
                        {row.map((col,x) => (
                            <div key={x} className="col" id={`${y}|${x}`} onClick={handleClick}>{getBoardElement(col)}</div>
                        ))}
                    </div>
                ))}
            </div>
        </div>
    )
}

export default Game