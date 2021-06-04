import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navbar from './components/Navbar'
//pages
import Create from './views/Create';
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
          <Route path="/create">
            <Create/>
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
