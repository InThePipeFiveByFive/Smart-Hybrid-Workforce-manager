//Source code provided by Lester Fernandez
//https://github.com/lesterfernandez
import { Navigate, Outlet, useLocation } from "react-router-dom";
import { useContext } from "react";
import { UserContext } from "../App";



const ProtectedRoute = () => {
    const { userData,setUserData } = useContext(UserContext);    
    const location = useLocation();
    console.log(userData);
    // const navigate = useNavigate();
    // console.log(userData);
    // console.log((userData==null));
    // console.log(sessionStorage.getItem("auth_data"));
    // if(userData==null){
    //     var auth_data = sessionStorage.getItem("auth_data");
    //     if(auth_data != null){
    //         setUserData(auth_data)
    //         for(var i = 0; i < 10000; i++){ 
    //         }
    //         return <Navigate to="/login" replace state={{from: location}}/>;
    //     }      
    // }
    if(userData && userData.expr_time > Date.now()){        //userData && userData.expr_time > Date.now()
        return <Outlet/>;
    }        
    else
        return <Navigate to="/login" replace state={{from: location}}/>
    // return (userData==null)?<Outlet/>: //Change ==
}

export default ProtectedRoute;