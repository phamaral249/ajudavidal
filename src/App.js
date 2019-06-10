import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import './App.css';
import Login from './components/Login'
import SignUp from './components/SignUp'
import Movies from './components/Movies'
import Error from './components/Error'

class App extends Component {


  render() {
    return (
      <BrowserRouter>
      	<Switch>
      		<Route path="/" component={Login} exact />
      		<Route path="/signup" component={SignUp} exact/>
      		<Route path="/movies" component={Movies} exact/>
      		<Route component={Error} />
      	</Switch>
      	
      </BrowserRouter>
    );
  }
}

export default App;
