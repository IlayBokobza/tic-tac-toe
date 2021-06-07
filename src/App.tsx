import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navbar from './components/Navbar'
//pages
import Create from './views/Create';
import Game from './views/Game';
import Home from './views/Home'

function App() {
  return (
    <Router>
      <div className="App">
        <Navbar/>
        <Switch>
          {/* views */}
          <Route exact path="/">
            <Home/>
          </Route>
          <Route exact path="/create">
            <Create/>
          </Route>
          <Route exact path="/gameroom/">
            <Game/>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
