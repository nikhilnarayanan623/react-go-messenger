import HomePage from "./components/home/HomePage";
import SideNav from "./components/partials/SideNav";
import Suggestions from "./components/home/Suggestions";

function App() {
  return (
    <div className='  flex  justify-center w-full'>
      <div className="w-2/12">
        <SideNav  />
      </div>
      <div className="w-7/12">
        <HomePage />
      </div>
      <div className=" w-3/12 p-5">
        <Suggestions />
      </div>
    </div>
  );
}

export default App;
