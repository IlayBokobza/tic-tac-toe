import { Link } from "react-router-dom"

const NavBar = () => {

    return (
        <div className="navbar">
            <h1>Tic Tac Toe</h1>
            <div className="links">
                <Link to="/">Home</Link>
                <button>Join Game</button>
                <Link to="/create">Create Game</Link>
            </div>
        </div>
    )
}

export default NavBar