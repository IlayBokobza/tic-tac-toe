import { HashRouter, Route, Switch } from 'react-router-dom';
import Navbar from './components/Navbar'
//pages
import Create from './views/Create';
import Game from './views/Game';
import Home from './views/Home'

function App() {
  const routes = [
    {
      path:'/',
      component:Home
    },
    {
      path:'/create',
      component:Create
    },
    {
      path:'/gameroom',
      component:Game
    },
  ]

  return (
    <HashRouter>
      <div className="App">
        <Navbar/>
        <Switch>
          {/* views */}
          {routes.map(route => (
            <Route exact path={route.path}>
              <route.component/>
            </Route>
          ))}
        </Switch>
      </div>
    </HashRouter>
  );
}

export default App;
