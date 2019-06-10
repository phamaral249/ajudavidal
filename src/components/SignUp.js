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

class SignUp extends Component {

constructor(props){
  super(props);
  
  this.state = {
    name:'',
    email:'',
    password:''

  };


 }


sendInfo() {

  const a = JSON.stringify(this.state)
  alert(a)
  axios.post('192.168.0.10:8080/user/', a)
        .then(res => {
            console.log(res)
            this.props.history.push('/');
        })
        .catch(error => {
            console.log(error.response)
        });


}  

render() {
    return (
      <div>
        <MuiThemeProvider>
          <div style={style_form}>
          <AppBar
             title="Sign Up"
           />
           <TextField
             required
             hintText="Username"
             floatingLabelText="Username"
             onChange={(event,newValue) => this.setState({name:newValue})}
             />
           <br/>
           <TextField
             required
             hintText="Email"
             floatingLabelText="Email"
             onChange={(event,newValue) => this.setState({email:newValue})}
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
             <RaisedButton label="Submit" primary={true} style={style_button} onClick={() => this.sendInfo()}/>
             <br/>

             <Link to="/">
                {"Already have an account? Log In"}
             </Link>

         </div>
         </MuiThemeProvider>
      </div>
    );
  }
}



export default SignUp;