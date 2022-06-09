import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
// import { useAppContext } from "../lib/contextLib";


function Login()
{
  const [identifier, setIdentifier] = useState("");
  const [secret, setSecret] = useState("");

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      console.log(identifier);
      console.log(secret);

      let res = await fetch("http://localhost:8100/api/user/login", 
      {
        method: "POST",
        body: JSON.stringify({
          id:null,
          secret:secret,
          active:null,
          FailedAttempts:null,
          LastAccessed:null,
          Identifier:identifier
        })
      });

      if(res.status === 200){
        alert("Successfully Logged In!");
        window.location.assign("./bookings");
      }else{
        alert("Failed to login");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  
  
  return (
    <div className='page-container'>
      <div className='content-login'>
        <div className='login-grid'>
          <div className='form-container-login'>
            <p className='form-header'><h1>WELCOME BACK</h1>Please enter your details.</p>
            
            <Form className='form' onSubmit={handleSubmit}>
              <Form.Group className='form-group' controlId="formBasicEmail">
                <Form.Label className='form-label'>Email<br></br></Form.Label>
                <Form.Control className='form-input' type="email" placeholder="Enter your email" value={identifier} onChange={(e) => setIdentifier(e.target.value)}/>
              </Form.Group>

              <Form.Group className='form-group' controlId="formBasicPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control className='form-input' type="password" placeholder="Enter your password" value={secret} onChange={(e) => setSecret(e.target.value)}/>
              </Form.Group>

              <Button className='button-submit' variant='primary' type='submit'>Sign In</Button>
            </Form>
            <p className='signup-prompt'>Don't have an account? <a className='signup-link' href='/signup'>Sign up for free!</a></p>
          </div>

          <div className='image-container'>
            <img className='login-image' src='https://i.pinimg.com/originals/43/90/d7/4390d72e6a6cb6086c73e570bb6c439d.jpg' alt='office'></img>
          </div>
        </div>

      </div>
    </div>
  )
}

export default Login