import { useState, useEffect } from "react";
import HomePage from "./components/home/HomePage";
import SideNav from "./components/partials/SideNav";
import Suggestions from "./components/home/Suggestions";
import { useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { selectIsLoggedIn } from "./features/slices/authSlice";

function App() {
  const loggedIn = useSelector(selectIsLoggedIn);
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(true);
  useEffect(() => {
    if (!loggedIn) {
      navigate("/sign-in");
    } else {
      setIsLoading(false);
    }
  }, [loggedIn, navigate]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
      <div className='  flex  justify-center w-full'>
        <div className='w-2/12'>
          <SideNav />
        </div>
        <div className='w-7/12'>
          <HomePage />
        </div>
        <div className=' w-3/12 p-5 '>
          <Suggestions />
        </div>
      </div>
  );
}

export default App;
