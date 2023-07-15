import { Outlet } from "react-router-dom";
import { useState, useEffect } from "react";
import axios from "axios";
import END_POINTS from "./constants/endpoints";

function App() {
  const [users, setUsers] = useState<any[] | null>(null);
  const fetchUser = async () => {
    try {
      const response = await axios.get(END_POINTS.GET_ALL_USERS, {
        headers: {
          authorization:
            "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUb2tlbklEIjoiY2FjYmFkYzQtNGY1OS00NzA5LTkxMjUtZjdkMDE0ZDg1YmE4IiwiVXNlcklEIjozLCJFeHBpcmVzQXQiOiIyMDIzLTA3LTE1VDAwOjI3OjQ4LjQ3ODIxODMrMDU6MzAiLCJSb2xlIjoidXNlciJ9.2npdjwPLg6QXEkyBIM3HnPs3LHlAqcBqQGC5tybNvWA",
        },
      });
      console.log(response);
      // setUsers(response)
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    fetchUser();
  }, []);
  return (
    <div>
      Home page
      <Outlet />
    </div>
  );
}

export default App;
