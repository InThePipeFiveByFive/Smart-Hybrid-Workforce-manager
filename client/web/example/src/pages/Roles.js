import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import Button from 'react-bootstrap/Button'
import { useState, useEffect, useContext } from 'react';
import RoleListItem from '../components/Role/RoleListItem';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../App.js';

function Roles()
{
  const [roles, SetRoles] = useState([]);

  const navigate = useNavigate();

  const {userData} = useContext(UserContext);

  //POST request
  const FetchRoles = () =>
  {
    fetch("http://localhost:8080/api/role/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        }).then((res) => res.json()).then(data => 
          {
            SetRoles(data);
          });
  }

  const AddRole = () =>
  {
    navigate("/role-create");
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchRoles()
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='resources-map'>
          {roles.length > 0 && (
            roles.map(role => 
            {
              return <RoleListItem id={role.id} name={'Role ' + role.role_name} lead={ role.role_lead_id } />
            }
          )
          )}
        </div>

        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddRole}>Add Role</Button>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Roles