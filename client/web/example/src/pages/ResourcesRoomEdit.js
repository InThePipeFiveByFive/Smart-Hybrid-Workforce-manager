import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../App.js'

const RoomEdit = () =>
{
  const [roomName, setRoomName] = useState("");
  const [roomLocation, setRoomLocation] = useState("");
  const [roomDimensions, setRoomDimensions] = useState("");

  const navigate = useNavigate();

  const {userData} = useContext(UserContext);

  //POST request
  const FetchRoom = () =>
  {
    fetch("http://localhost:8080/api/resource/room/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            id: window.sessionStorage.getItem("RoomID")
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        }).then((res) => res.json()).then(data => 
          {
            setRoomName(data[0].name);
            setRoomLocation(data[0].location);
            setRoomDimensions(data[0].dimension);
          });
  }

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/resource/room/create", 
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("RoomID"),
          building_id: window.sessionStorage.getItem("BuildingID"),
          name: roomName,
          location: roomLocation,
          dimension: roomDimensions
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if(res.status === 200)
      {
        alert("Room Successfully Updated!");
        navigate("/resources");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    FetchRoom();
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT ROOM</h1>Please update the room details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Name<br></br></Form.Label>
              <Form.Control name="bName" className='form-input' type="text" placeholder="Name" value={roomName} onChange={(e) => setRoomName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Location<br></br></Form.Label>
              <Form.Control name="bLocation" className='form-input' type="text" placeholder="Location" value={roomLocation} onChange={(e) => setRoomLocation(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Dimensions<br></br></Form.Label>
              <Form.Control name="bDimensions" className='form-input' type="text" placeholder="10x10" value={roomDimensions} onChange={(e) => setRoomDimensions(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Room</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default RoomEdit