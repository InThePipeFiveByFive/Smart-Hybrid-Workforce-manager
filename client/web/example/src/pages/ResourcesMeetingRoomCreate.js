import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useContext, useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"
import { UserContext } from '../App.js'

const CreateMeetingRoom = () =>
{
  const [meetingRoomName, SetmeetingRoomName] = useState("");
  const [meetingRoomLocation, SetmeetingRoomLocation] = useState("");
  const [meetingRoomCapacity, SetmeetingRoomCapacity] = useState("");

  const navigate = useNavigate();
  const {userData} = useContext(UserContext);

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/resource/create", 
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
          id: null,
          room_id: window.sessionStorage.getItem("RoomID"),
          name: meetingRoomName,
          location: meetingRoomLocation,
          role_id: null,
          resource_type: 'MEETINGROOM',
          decorations: `{
            capacity: ${meetingRoomCapacity}
          }`
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if(res.status === 200)
      {
        alert("Meeting Room Successfully Created!");
        navigate("/resources");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>CREATE YOUR MEETING ROOM</h1>Please enter your meeting room details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={meetingRoomName} onChange={(e) => SetmeetingRoomName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Location<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Location" value={meetingRoomLocation} onChange={(e) => SetmeetingRoomLocation(e.target.value)} />
            </Form.Group>
                      
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Capacity<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Capacity" value={meetingRoomCapacity} onChange={(e) => SetmeetingRoomCapacity(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Meeting Room</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateMeetingRoom