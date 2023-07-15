import { Card, CardHeader, Typography, Avatar } from "@material-tailwind/react";
import { AiOutlineHeart } from "react-icons/ai";
import { BiMessageRounded } from "react-icons/bi";
import { BsThreeDots } from "react-icons/bs";
import {FiBookmark} from 'react-icons/fi'
import {IoPaperPlaneOutline} from "react-icons/io5"

export default function PostCard() {
  return (
    <div className=' mt-3 pt-5 pb-10 border-b border-gray-300'>
      <div className='flex items-center justify-between'>
        <div className='flex items-center mb-1'>
          <Avatar
            size='md'
            src='data:image/jpeg;base64,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'
          />
          <Typography
            variant='h5'
            color='blue-gray'
            className='mb-2 ml-2 mt-2 font-medium leading-[1.5]'
          >
            Abin_abin <span className='text-sm font-semibold'>.6h</span>
          </Typography>
        </div>
        <div className='p-2 cursor-pointer '>
          <BsThreeDots />
        </div>
      </div>
      <Card
        shadow={false}
        className='relative rounded-md  grid h-[38rem] w-[28rem] items-end justify-center overflow-hidden text-center'
      >
        <CardHeader
          floated={false}
          shadow={false}
          color='transparent'
          className="absolute inset-0 m-0 h-full w-full rounded-none bg-[url('https://images.unsplash.com/photo-1552960562-daf630e9278b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80')] bg-cover bg-center"
        >
          <div className='to-bg-black-10 absolute inset-0 h-full w-full bg-gradient-to-t from-black/80 via-black/50' />
        </CardHeader>
      </Card>
      <div className='flex pt-2 '>
        <div className='w-11/12 gap-3 flex'>
          <AiOutlineHeart className='w-7 h-7 cursor-pointer hover:bg-red-500' />
          <BiMessageRounded className='w-7 h-7 cursor-pointer' />
          <IoPaperPlaneOutline className='w-7 h-7 cursor-pointer' />
        </div>
        <div className='w-1/12'>
          <FiBookmark className='w-7 h-7 cursor-pointer' />
        </div>
      </div>
    </div>
  );
}
