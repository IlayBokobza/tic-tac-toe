import { useState } from "react"

type props = {
    title:string,
    text:string,
    onlyAccept:boolean,
    acceptText?:string,
    cancelText?:string,
    acceptFunc:Function,
    closeFunc:Function,
    showInput?:boolean,
}

const Popup = ({title,text,onlyAccept,acceptText,acceptFunc,closeFunc,showInput,cancelText}:props) => {
    const [bgStyle,setBgStyle] = useState('')
    const [boxStyle,setBoxStyle] = useState('')

    const closePopup = () => {
        setBoxStyle('re-fade-from-top 1s forwards')
        setBgStyle('re-popup-bg-fade 1s forwards')
        setTimeout(closeFunc,1000)
    }

    return (
        <div style={{animation:bgStyle}} className="popup">
            <div style={{animation:boxStyle}} className="content">
                <h2>{title}</h2>
                <p>{text}</p>
                {showInput && <input type="text" id="popup-id"/>}
                <div className="action-box">
                    <button onClick={() => {
                        const input = document.getElementById('popup-id') as HTMLInputElement
                        const inputVal = (input) ? input.value : null
                        acceptFunc(inputVal)
                        closePopup()
                    }}>{acceptText||'Ok'}</button>
                    {!onlyAccept && <button onClick={() => {closePopup()}}>{cancelText||'Cancel'}</button>}
                </div>
            </div>
        </div>        
    )
}

export default Popup