import React from "react";
import Avatar from "./Avatars";
import PostCard from "./PostCard";
import { IoIosArrowDropright } from "react-icons/io";

const HomePage: React.FC = () => {
  return (
    <div className='pt-4'>
      <div className='p-2 w-11/12 items-center ml-10  flex flex-wrap justify-center '>
        <Avatar />
        <Avatar />
        <Avatar />
        <Avatar />
        <Avatar />
        <Avatar />
        <Avatar />
        <IoIosArrowDropright className='w-7 h-7 cursor-pointer text-blue-gray-700' />
      </div>
      <div className='flex flex-col items-center '>
        <PostCard />
        <PostCard />
        <PostCard />
        <PostCard />
      </div>
    </div>
  );
};

export default HomePage;
