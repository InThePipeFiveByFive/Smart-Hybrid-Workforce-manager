import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect, useCallback, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { UserContext } from '../App.js'

const TeamPermissions = () =>
{
  const [teamName, setTeamName] = useState(window.sessionStorage.getItem("TeamName"));
  //const [teamPermissions, SetTeamPermissions] = useState([]);

  const {userData} = useContext(UserContext);

  const [createTeamIdentifier, SetCreateTeamIdentifier] = useState("") // allows a user to update the team
  const [createTeamIdentifierId, SetCreateTeamIdentifierId] = useState("")
  const [viewTeamIdentifier, SetViewTeamIdentifier] = useState("") // allows a user to view the team
  const [viewTeamIdentifierId, SetViewTeamIdentifierId] = useState("")
  const [deleteTeamIdentifier, SetDeleteTeamIdentifier] = useState("") // allows a user to delete the team
  const [deleteTeamIdentifierId, SetDeleteTeamIdentifierId] = useState("")


  const [createTeamUser, SetCreateTeamUser] = useState("") // allows users to add a user of the team
  const [createTeamUserId, SetCreateTeamUserId] = useState("")
  const [viewTeamUser, SetViewTeamUser] = useState("") // allows users to view users of the team
  const [viewTeamUserId, SetViewTeamUserId] = useState("")
  const [deleteTeamUser, SetDeleteTeamUser] = useState("") // allows users to remove users of the team
  const [deleteTeamUserId, SetDeleteTeamUserId] = useState("")

  let handleSubmit = async (e) => {
    e.preventDefault();
    if (createTeamIdentifier && createTeamIdentifierId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "CREATE", "TEAM", "IDENTIFIER");
    }
    if (!createTeamIdentifier && createTeamIdentifierId != null)
    {
      RemovePermission(createTeamIdentifierId);
    }

    if (viewTeamIdentifier && viewTeamIdentifierId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "VIEW", "TEAM", "IDENTIFIER");
    }
    if (!viewTeamIdentifier && viewTeamIdentifierId != null)
    {
      RemovePermission(viewTeamIdentifierId);
    }

    if (deleteTeamIdentifier && deleteTeamIdentifierId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "DELETE", "TEAM", "IDENTIFIER");
    }
    if (!deleteTeamIdentifier && deleteTeamIdentifierId != null)
    {
      RemovePermission(deleteTeamIdentifierId);
    }

    if (createTeamUser && createTeamUserId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "CREATE", "TEAM", "USER");
    }
    if (!createTeamUser && createTeamUserId != null)
    {
      RemovePermission(createTeamUserId);
    }

    if (viewTeamUser && viewTeamUserId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "VIEW", "TEAM", "USER");
    }
    if (!viewTeamUser && viewTeamUserId != null)
    {
      RemovePermission(viewTeamUserId);
    }

    if (deleteTeamUser && deleteTeamUserId === "")
    {
      AddPermission(window.sessionStorage.getItem("TeamID"), "DELETE", "TEAM", "USER");
    }
    if (!deleteTeamUser && deleteTeamUserId != null)
    {
      RemovePermission(deleteTeamUserId);
    }
    
    alert("Permissions successfully updated.")
  };

  async function AddPermission(id, type, category, tenant) 
  {
    try
    {
      let res = await fetch("http://localhost:8080/api/permission/create", 
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
          permission_id: id,
          permission_id_type: "TEAM",
          permission_type: type,
          permission_category: category,
          permission_tenant: tenant
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if (res.status === 200) { 
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

  async function RemovePermission(id) 
  {
    try
    {
      let res = await fetch("http://localhost:8080/api/permission/remove", 
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
            id: id,
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if (res.status === 200) { 
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

  //POST request
  const FetchTeamPermissions = useCallback(() =>
  {
    fetch("http://localhost:8080/api/permission/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            permission_id: window.sessionStorage.getItem("TeamID"),
            permission_id_type: "TEAM",
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        }).then((res) => res.json()).then(data => 
        {
          //SetTeamPermissions(data);
          data.forEach(setPermissionStates);
        });
  },[]);

  function setPermissionStates(permission)
  {
    if (permission.permission_type === "CREATE")
    {
      if (permission.permission_tenant === "IDENTIFIER")
      {
        SetCreateTeamIdentifier(true);
        SetCreateTeamIdentifierId(permission.id);
      }        
      if (permission.permission_tenant === "USER")
      {
        SetCreateTeamUser(true);
        SetCreateTeamUserId(permission.id);
      }
    }
    if (permission.permission_type === "VIEW")
    {
      if (permission.permission_tenant === "IDENTIFIER")
      {
        SetViewTeamIdentifier(true);
        SetViewTeamIdentifierId(permission.id);
      }        
      if (permission.permission_tenant === "USER")
      {
        SetViewTeamUser(true);
        SetViewTeamUserId(permission.id);
      }
    }
    if (permission.permission_type === "DELETE")
    {
      if (permission.permission_tenant === "IDENTIFIER")
      {
        SetDeleteTeamIdentifier(true);
        SetDeleteTeamIdentifierId(permission.id);
      }        
      if (permission.permission_tenant === "USER")
      {
        SetDeleteTeamUser(true);
        SetDeleteTeamUserId(permission.id);
      }
    }
  }
    
  useEffect(() =>
  {
    FetchTeamPermissions();

    setTeamName(window.sessionStorage.getItem("TeamName"));
  }, [setTeamName, FetchTeamPermissions])
    
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>Team {teamName}</h1> Configure Permissions</p>

          <Form className='form' onSubmit={handleSubmit}>
            <h2 className='permission-category'>Team</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the team to update team information  
                <input className='checkbox' type="checkbox" checked={createTeamIdentifier} onChange={(e) => SetCreateTeamIdentifier(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the team to view team information and enrollment  
                <input className='checkbox' type="checkbox" checked={viewTeamIdentifier} onChange={(e) => SetViewTeamIdentifier(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the team to delete the team  
                <input className='checkbox' type="checkbox" checked={deleteTeamIdentifier} onChange={(e) => SetDeleteTeamIdentifier(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Users</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users to add users to the team
                <input className='checkbox' type="checkbox" checked={createTeamUser} onChange={(e) => SetCreateTeamUser(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users to view users of the team  
                <input className='checkbox' type="checkbox" checked={viewTeamUser} onChange={(e) => SetViewTeamUser(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users to remove users from the team
                <input className='checkbox' type="checkbox" checked={deleteTeamUser} onChange={(e) => SetDeleteTeamUser(e.target.checked)}/>
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Team Permissions</Button>
          </Form>
          
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default TeamPermissions