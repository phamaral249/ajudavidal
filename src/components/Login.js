import React, { Component } from 'react';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import AppBar from 'material-ui/AppBar';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import { Link } from 'react-router-dom'
import axios from 'axios';


const style_button = {
 margin: 15,
};

const style_form = {
  textAlign: 'center',

}

class Login extends Component {

constructor(props){
  super(props);
  
  this.state = {
    username:'',
    password:'',
  }

 }


sendLogin() {

  const a = JSON.stringify(this.state)
  alert(a)
  axios.post('localhost:8080/login/', {a})
        .then(response => {
            console.log(response)
            this.props.history.push('/movies');
        })
        .catch(error => {
            console.log(error.response)
        })

}  

render() {
    return (
      <div>
        <MuiThemeProvider>
          <div style={style_form}>
          <AppBar
             title="Login"
           />
           <TextField
             required
             hintText="Username"
             floatingLabelText="Username"
             onChange={(event,newValue)=> this.setState({username:newValue})}
             />
           <br/>
             <TextField
               required
               type="password"
               hintText="Enter your Password"
               floatingLabelText="Password"
               onChange={(event,newValue) => this.setState({password:newValue})}
               />

             <br/>
             <RaisedButton label="Submit" primary={true} style={style_button} onClick={(event) => this.sendLogin()}/>
             <br/>

             <Link to="/signup"> {"Don't have an account? Sign Up"}
             </Link>

         </div>
         </MuiThemeProvider>
      </div>
    );
  }
}



export default Login;