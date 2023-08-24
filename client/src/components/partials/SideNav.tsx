import React from "react";
import {
  Card,
  Typography,
  List,
  ListItem,
  ListItemPrefix,
  ListItemSuffix,
  Chip,
} from "@material-tailwind/react";
import { Link } from "react-router-dom";
import { PowerIcon } from "@heroicons/react/24/solid";
import { BiSolidHome, BiSearch } from "react-icons/bi";
import { MdOutlineExplore } from "react-icons/md";
import { BsMessenger } from "react-icons/bs";
import { useDispatch } from "react-redux";
import { clearToken } from "../../features/slices/authSlice";

const SideNav: React.FC = () => {
  const dispatch = useDispatch();
  const handleLogout = () => {
    dispatch(clearToken());
    location.reload();
  };
  return (
    <Card className='fixed top-0  h-full w-full max-w-[17rem] p-4 rounded-none'>
      <div className='mb-2 flex items-center gap-4 p-4'>
        <Link to='/'>
          <img
            className='h-10 w-10'
            src='data:image/jpeg;base64,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'
            alt='brand'
          />
        </Link>
        <Typography variant='h5' color='blue-gray'>
          Messenger
        </Typography>
      </div>
      <List>
        <hr className='my-2 border-blue-gray-50' />
        <Link to='/'>
          <ListItem>
            <ListItemPrefix>
              <BiSolidHome className='h-7 w-7' />
            </ListItemPrefix>
            <span className='mt-2'>Home</span>
          </ListItem>
        </Link>
        <Link to='/search'>
          <ListItem>
            <ListItemPrefix>
              <BiSearch className='h-7 w-7' />
            </ListItemPrefix>
            <span className='mt-1'>Search</span>
          </ListItem>
        </Link>
        <Link to='/explore'>
          <ListItem>
            <ListItemPrefix>
              <MdOutlineExplore className='h-7 w-7' />
            </ListItemPrefix>
            <span className='mt-1'>Explore</span>
          </ListItem>
        </Link>
        <Link to='/chats'>
          <ListItem>
            <ListItemPrefix>
              <BsMessenger className='h-6 w-6' />
            </ListItemPrefix>
            <span className='mt-1'>Messages</span>
            <ListItemSuffix>
              <Chip
                value='14'
                size='sm'
                variant='ghost'
                color='blue-gray'
                className='rounded-full'
              />
            </ListItemSuffix>
          </ListItem>
        </Link>
        {/* <Link to='/sign-in' onClick={handleLogout}> */}
          <ListItem onClick={handleLogout}>
            <ListItemPrefix>
              <PowerIcon className='h-7 w-7 ' />
            </ListItemPrefix>
            <span className='mt-1'> Log Out</span>
          </ListItem>
        {/* </Link> */}
      </List>
    </Card>
  );
};

export default SideNav;
