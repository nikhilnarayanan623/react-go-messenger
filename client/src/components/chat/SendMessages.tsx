import React, { useState, useRef, useEffect } from "react";
import {
  ListItem,
  ListItemPrefix,
  Avatar,
  Typography,
} from "@material-tailwind/react";
import UserChat from "../../api/chat/chats";
import { useParams } from "react-router-dom";
import { Message } from "../../types/Chat";
import { useSelector } from "react-redux";
import { selectCurrentChat } from "../../features/slices/chatSlice";
import { toast } from "react-toastify";
import { useWebSocketContext } from "../../features/contexts/socketContext";

const SendMessages: React.FC = () => {
  const userChat = UserChat();
  const { chatId, receiverId } = useParams();
  const divRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const [recentMessages, setRecentMessages] = useState<
    Record<string, Message[]>
  >({});
  const [message, setMessage] = useState<string>("");
  const currentChat = useSelector(selectCurrentChat);
  const { socket } = useWebSocketContext();

  const fetchRecentMessage = async () => {
    try {
      const response = await userChat.getRecentMessages(chatId ?? "");
      const updatedMessages: Record<string, Message[]> = {};
      updatedMessages[chatId ?? ""] = response?.data.reverse();
      setRecentMessages((prevMessages) => ({
        ...prevMessages,
        ...updatedMessages,
      }));
    } catch (error: any) {
      toast.error(error?.data?.error[0] || "Something went wrong", {
        position: toast.POSITION.BOTTOM_RIGHT,
      });
    }
  };

  const handleTypeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const message = e.target.value;
    setMessage(message);
  };

  const handleSendMessage = async () => {
    try {
      const trimmedMessage = message.trim();
      if (trimmedMessage === "") {
        return;
      }
      const response = await userChat.sendMessage(chatId ?? "", trimmedMessage);
      const updatedMessage: Message = {
        receiver_id: parseInt(receiverId ?? ""),
        chat_id: parseInt(chatId ?? ""),
        is_current_user: true,
        content: trimmedMessage,
        created_at: Date.now(),
      };
      socket?.send(JSON.stringify(updatedMessage));

      setRecentMessages((prevMessages) => ({
        ...prevMessages,
        [chatId ?? ""]: [...(prevMessages[chatId ?? ""] || []), updatedMessage],
      }));

      if (response?.success) {
        setMessage("");
      }
    } catch (error: any) {
      toast.error(error?.data?.error[0] || "Something went wrong", {
        position: toast.POSITION.BOTTOM_RIGHT,
      });
    }
  };

  useEffect(() => {
    if (socket) {
      const handleMessage = (event: any) => {
        try {
          const messageData = JSON.parse(event.data);

          // Check if the received message belongs to the current chat
          if (messageData.chat_id === parseInt(chatId ?? "")) {
            const message: Message = {
              chat_id: parseInt(chatId ?? ""),
              receiver_id: parseInt(receiverId ?? ""),
              is_current_user: false,
              content: messageData.content,
              created_at: Date.now(),
            };
            setRecentMessages((prevMessages) => {
              const chatId = messageData.chat_id.toString();
              const updatedMessages = {
                ...prevMessages,
                [chatId]: [...(prevMessages[chatId] || []), message],
              };
              return updatedMessages;
            });
          }
        } catch (error) {
          console.log("Error parsing incoming message:", error);
        }
      };

      // Set the onmessage event handler
      socket.onmessage = handleMessage;

      // Clean up the event handler when the component unmounts or chatId changes
      return () => {
        socket.onmessage = null;
      };
    }
  }, [socket, chatId, receiverId]);

  useEffect(() => {
    scrollDiv();
  }, [recentMessages]);

  useEffect(() => {
    fetchRecentMessage();
  }, [chatId]);

  useEffect(() => {
    inputRef && inputRef.current?.focus();
  }, [inputRef]);

  function handleKeyPress(event: React.KeyboardEvent<HTMLInputElement>) {
    if (event.key === "Enter") {
      handleSendMessage();
    }
  }
  const scrollDiv = () => {
    const divElement = divRef.current;
    if (divElement) {
      divElement.scrollIntoView({ behavior: "smooth" });
    }
  };

  const profilePic = false;

  return (
    <div className='w-full h-full flex flex-col'>
      <div>
        <ListItem
          className='border-b-2 rounded-none focus:bg-transparent w-full p-4 mt-0.5 hover:bg-transparent active:bg-transparent focus:ring-0 '
          ripple={false}
        >
          <ListItemPrefix>
            <Avatar
              variant='circular'
              alt='candice'
              src={
                currentChat?.profile_picture
                  ? currentChat?.profile_picture
                  : "data:image/jpeg;base64,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"
              }
            />
          </ListItemPrefix>
          <div>
            <Typography variant='h6' color='blue-gray'>
              {currentChat?.first_name}
            </Typography>
          </div>
        </ListItem>
      </div>

      <div
        className='flex-grow overflow-y-auto'
        style={{
          scrollbarWidth: "thin",
          scrollbarColor: "black",
        }}
      >
        {recentMessages[chatId ?? ""]?.map(
          ({ content, chat_id, is_current_user, receiver_id }, index) => {
              return (
                <ListItem
                  ref={divRef}
                  className={`rounded-none active:bg-transparent focus:bg-transparent focus:ring-0 hover:bg-transparent w-full pt-4 ${
                    is_current_user ? "flex-row-reverse" : "flex-row"
                  }`}
                  ripple={false}
                  key={index}
                >
                  <ListItemPrefix>
                    <Avatar
                      variant='circular'
                      alt='candice'
                      size='sm'
                      src={
                        is_current_user
                          ? profilePic
                            ? ""
                            : "data:image/jpeg;base64,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"
                          : currentChat?.profile_picture
                          ? currentChat.profile_picture
                          : "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAHYA7AMBIgACEQEDEQH/xAAcAAABBQEBAQAAAAAAAAAAAAADAAECBAUGBwj/xABGEAACAQMCAgcGAgcECAcAAAABAgMABBEFEiExBhMiQVFhcRQygZGhwVLRByMzQmKx8CRTcoIVFjREY5Oi4SVDVXOSsvH/xAAaAQACAwEBAAAAAAAAAAAAAAABAwACBAUG/8QAKBEAAgIBBAEEAQUBAAAAAAAAAAECEQMEEiExQRMiMlFhBVJxkcEU/9oADAMBAAIRAxEAPwDPpVNSh4OufMVJ4OyGiO5fqK3uaXDMW1tcAaanwRSFXRQbFLFSxTUSDEA8xkeFWLe8uLbAjk3J+CTiPh3igVb03T7jUZikAAVP2kj8Fj9fyrNqFi2XlqhmKWRS9hftdYidlSRZInJwuO0GPliukttOlxvupWRfDdx+JqjZR2Wkf7IgmuMYa4lHH4Du+FXbR2kuI2unLzPnq4z+6O9jXkNVODk3hTS/J24PLt95qRRJEu1AQPM5olCkl2yQxjj1hI9AATmpySLFGXc4VRk1zHbYrsaaVIYzJIcKPrUbWYXEIlC7dxPCub1XUOumAaRY1OduWA2jx9au6XrGmQWYSbULWMqxGGmXNaHp5KF1YxxUYm5SrM/1i0T/ANWsv+etTXXdIcgJqlmc/wDHX86T6WT9r/oXZekcIjOxwqjJPhVexuTdda/7gbaorE1jV47ldlpPG0I4s6uDy7x41d0T+yabaQ/vyseDHj55+VNeFxx2+y+322bFKg3E6QmIMeMjhFpXc/s9s8vLby+eKRtbK1ZDULRb21khB2OUISQc0Pca81cSq7LM8nWKxVgZGOCOdehXdz7LfI3/AJTqNw+9c30nsQNUNzGyiO5QSZz3jgft867v6NlWPI4T6f8AgjUwm4ppmAxfYydZLsYYZesbB+GaPa2UlxgjgnianFHCXwoaVu4d1aCRTsO04QeC13M+qWPjGqYjDp3PmbJw28NsucDPeTzoqNv8cedDW3XPbLP6mjelcuUnJ2zoqKiqQqVReRE95gKryX0aghOJoBA6kqbg2e34VRqUjl2LMeJqNWABpAkHgSKltpsV6NpUcFOgiS57Mi7h9aJ7Okn7JwG/C1V6NDOF4SoGH1FZ5wlBXjY+E4y9syDwSx+8h9aGR/RrTjlBP6uXh+FxVlUSQdtEOe8VmWuceJxH/wDIpfGRn6bZe1yFpGKW6HtsBxJ/Cvn591bD31nBEtvHLBDCnKJXBOfE95PnVFtLg6sIGkIHIM5I+VVZbB4SqxYYtyRf6+tY8sY6qW7JOkvA6EXgXtjbNF9VskXe8/ZHEnYx+1b+jR9TBJqF4dry8ADzVfD1rlLeyjhkD3BDtGQTtOQG7lA7z3/LzrH6e69qkc0FlExtoJIgQEPaIyRjy7uXjzrnZNPHLP08T4+zU5z2XPg72z1WLUtcMcLoVtUIcA5KlvHwOBy86Fr2qxxxSyytttoFLEgZ3Y/rhXJfo/t5LPSrlmbDTzHPpgDHzzVW/wBYOp6t7NbHNnCSpPdK+QCR5AZ+fpQholLPtj0iOSjDcVrfTNa6S3s17BptzMzcFUJhY07hxwPX1q5J0N16AqJrERFhwDSJ9jXof6Obov1sDybnRSVz+HIx8q1Nek3X+0H9moH3rvRhGPCMbm27PJv9UtZ/uYv+aKG/RfWV/wBy3eki/nXfXGodVK8MFvJcPGAZNhVVTPIEkjj5VG3vLq7mTq7N7eFf2rXAw2fwqAePrnFW2oG5nBWWhXsWo23ttm8Me/JZsYOOOMj0r0HT3Mt9CDyRcKPhVXUD7Te9UP2dvFvlx5/mB8iao3WrroWk3GoMoMqQhIUPfIeXwHEn0rj/AKhjcpKvPBrxSWx2XrvVYL7Vri3t5AzWJCMo/EeJP2+Bq70luVXTYgWCid0GSccyB968r6FTSQap107s3tZKuzHixPHcfPP866jp/qQjtdHhzgmaMnjyCnJP8qyz0ezNDGi0WnFP6NjpBcXCnTZbftZjKvGf3vy5Hj50G8dL2ygRiR1co255rngVPxwfhRbx98cJH7jAA+ROPvVa7Uqjui5OAWA78cjTcKqMWlTTDKPafTJwQRwriMfHvNFpmZV5kAeZoL3cK83B9K0NuTtgSUVSLFAubhYF8WPIUhdRbQ27ArLnfrJS3dnhQSCRkkMjEsTx5Co0qVOhilJ0kLlkjHtiptwp8UsVux6JdzMWTVu6gPtpitHK1ErXQ3GEAVpsUbbTbaFhBYOacMynIJFTIpbarKMZdotGTj0wsd5Mg4Nn1pJcPGrlGfew/EDk+ZIPChbaW2sWTQwfxNcNXJcMs2bJ2TPNKZcYJzgA9+KyOlsMftdnIMswjfBZie8Vex3CnudCvNVgt7i2nt2QI/ZLHK4wcZ8xxrI8LwS3S6NKzLJGomRcai0HR+PT7Rj185brXH7iknv8cZ+FVNKjSO5ijUYULjFWp9A1W1G6WwlCnjlMOOPoc1QJktpe2hVsHKuCpxTsajHlAnua5PR+gDCDpDDGOCtEyDB58M/atLXNXjS2vdRiZWCsQueW7OAD8a830vX7vTp4p7RYy0eQhbtDljy8as6d0mu7F22xRur81YHn/wDKn2hNM2bGTSS0sN/exzxJMZcscrcFlAyccyCDw9K39DSSPTIhKrrzKI3NEJO0HzAxXMDpzP32MZ/zY+9Rn6bTSRFUtBGx5tvyQKNolM3tYvYbSO4htwDc3Rw/kSAo+2BXIdK3F6jrn9SAVjxyxlePxP0qtPqrSjGwgkgk541Tnuzc8FT3jyB3eHAfIUqUE5qX0XTe1xBIWjKmPgUwV9Ryq70mlGq65p5GepihRv8AM5Bx9BQorG8mO2K0uHP8MTH7VpaZ0Z1SW5j32xt4RJl5Z2AA28Twzny5UqajuUmxkd1UdDOwt4WRj+qyGTj7uCDj08KBc3rFisXAA8/GqtxGySyQTEMUYqRnI4ciPI86jWSGJ3wh0ppLliLE8yT6mm9KfFLFbIaLJLvgyz1cI9cjUqkFp9lbsekhj57Mc9TOf8A8VMLUttSAp/XQnsjiltomKkBVbCPimIouKiRVbDQLFNtouKbFGyUBIpsUXFCmligTfNIqL51LJ/A4FNIUiQySsqIoyWY4ArB1LpN1WUsbZ3P95IpAHoO/6VkGy1vWGDzRzMp4gy9hR6CqPIvAxY/L4Ne96RwLItvYL10kjhFc8FBJx8edehQw/wCjbyG3IVIXBjjBbJbCsM/9Q+VeZ2vRO5UrLJdIHQhlWNSxyOI48O8V6Zqbe06hDNHxSRI5EPjjeQB8do+IrBrIynCn5N2lUKkokNdvLy3W16iQpDJCPdHf38fSuTOp6re2E+oLZxyQW8/UAvPiQybtoAAThknxrvLizXU9JhUEKxiR42PccCuZt7KTTU1Wwv4rqO2vzHPFPBA0wguYzkEheJU4XP8Ah8650IL1mp9Prk6uScvQXp9rsyDL1N3Gmt6VbwySt1azMqOrN3K2QOPmceFdFa6f0euMLPZ21vLy2tGoyfI4HHyPGp9HNBsZr17vWI4Gics0kUNrP+vcgrudpB3Bm4DvPOtI6a1skCh+uaNtscuO1wPZJPjjGfjWmMXGPZlg9zqgQ6KaQQD7LHx8U/71TutG0GyBE9lGJB7qmAHcB3gnh9a6kA4AOM9+OVZ0qyzIsoJ/WttUfhXu+fP5Va39jNq+jidQks4buC2t9HimuJzmK3jhUEKObNwyR64HrR9PfXLq7kgsdIshcRxdaoa5LFkzgkEJjgSO7vFdBcdHbeSCaZrhl1IziSKXqmZCgUr1bAHOME8u/jVCzjPRqyvZLKELqktubazgs7SVYYd7Dc5dgMnIB8tuBmhsVpvryIcppvaUNH6TX11ZR3CHqwxI2FVIODjuHEVv3U0tzDFbKVFxJC8roeZ7XEAf5cehNZnRvQyBboYmjtYAAA4wWxyFXySL2O4DDcJ2UeiyAfXc/wA6waeMpT3v43wb9SlUI+fLMvVIlS6TbnBgjPHn7uPtVTbWnrkiz6rOyYxHiLI7yOf1NUdtemwKsatcnmtS7yyoGFp9tTxUsU6zODC04XhRAtPtqWQHtp9tE20ttBhQPFSAqWKWKoyxLFNVObWtPj4CUyH/AIa5+vKqcnSSLOyK3kZv4mAx60tzihihJ+DXINUrrUrS1JWSQM45onaP/asG71K6usiSTah/dTgKpAkrtiGByyeVUeT6GrD9mvNrNxO/V2sYRm5cNzflUoNHuJ3Et5KVJ7veb58hWdb3U9su2CTbnm20ZPxxRDqN2f8AeZPnVVJP5F3Br4nQ21hbW3GKIbvxHifmatBa5A390ynbcTSeSEk/Sqb6Vr2pvloZljPITSbQB5gn7UxZPCQl4n5Z2Nzf2dtxuLqFD5vxrQ0fUYdd0iVNNYy3Fk5VVI2lkbBBGfA//WuLsuhEg/2q6VB3pAvH5n8q6fQ9Ng0S5R9MTbM3vvISxde8Hy/kSKXnuUG5cUN00o450ubPQAMADvxT8eFZNldTXU80Tu0J4OojIbI5HmO4j61e9lUjDyzOP4pD9sVzU01Z2OSxQBboGDIXD8928n4canFEkQ2oMDzJNTokJKvYL9ykDFBto+pt4oiQxRApOOfCrCsOoYZ47h96HUArHphw5cKR8PGgeyxAnY0i/wCGRvzqFkWPhmsC6YRXF1cSqwWFx1Q24Ejc1UfE7jWhes1haS3PtEjCNc7JACD4DkD9a5+6vrnUBE1ysaBFHYQkjcRxPGnYcbnIzajMsUfyVACRkncTxJ8SeZ+dLbRNtLFdbo4QPbT4qe2n20SEQKephacLUshDFLFE204WgQFtp9tGCU+ygw2edsxBAX3j9KfhGMA/manZ20lxcR28OGlkP/6fTFdZZdHLODD3ANxJ4t7o9B+dYlGzfKaicpb2txdtthheUeCDh8TyrUg6PXjgGVo4h4Z3EfLhXXIiqoVVCqOQAwBT7aasaEPNLwc9F0bt1I66aWT/AA4UH7/WrkWk2UPu2yMfFxu/nWmVpitXSihblJ9sAiBfdUADuAxUsUTbinxVrKUCxx40EiQSOx/VxKnFgePeTjw7qubaFdFVQI/AOeJ8AOf5epFLy04PdyXx7lJbQlmTZxW0yL2oh21HeD7w/rwro0knnVWjCRRsMqzdokeg4D5muPu9Vg08y3GpXCwx4UInM57RwPE4xVvoj0osb/r7W2kkMcWDGWjIxnPZA+BxXGxKabTR3Yzi0qZ1caFFwXdyf3mx9qlQ4mkdtzpsXHBSePqfD0oc100cmxLaeVgMkonZHxPCmlytd+1RzN1cz7X9zgDt9OFXbdZFhVZmLSAdpjzqub2bl7FIv+J0H3qQubjmbJ9neVkUn5USWWiMgjjxGOFA6qZB+quWPlKu8fY/WiOvWxbSXXcO4lSKyOkGpzaLp806OJ3WNmRX4MuBzOOY+XdUSt0BukVtcu5JyLUhRtbB2nILY4/IZ+J8qohSazdC1e01bswSs0kUYLI/Fu1xJJ7/AAPmK2AvGt+ki4wbfbONrMm+YIqabbR9lLZWmzKB20+DRdtNig5BogBUgtPipAVNxKIhakFqYWphaO4FEAtOFooWpbKNko4/oLYObeXVLlcSTnZF/DGPD1P8hXU4poIUggjhiUBI1CqPACpVnTofLl2NimxT0qtYKIkU2KlTVLBQ1LFSpwKlk2jAUM26tcCdiWKrhUPujz9fyo+2n21O+yJM4j9JCg2UXj1yH/pcfasToLq9npt3cw3plAujEkcijOxgW4k93vDj/R1P0i6hZ3MFtBbTCSVJT1m3kBjx7+OK4YjskUhxTv8AJrxycaPojTb8yN7Nc4E4HZYe7IPEefiK0GXcCG5EYrxDQul95HdQxatO01mcAvtAeI9zAjw7/KvSdH6VW10xhjuIrwr+/Acn4rz+I4VmnBpnQhlUzb/0XZ/3IqxDBFACsKBFJ7qpDWbb8M2f/bNQl1cbcW8MjNjm6lVHqTVKY0uXl3Haxb37TH3UHNj4VwvTXVLeDR7+O7uYxf3URRIgckDuAHh51Q1bp3a2103URvfzqSGlztjTyXx9a89vbmW+uprqc5lmcuxJ+lOx4/LM2XMkqR1f6NF/t18/cI1HzJr0AV5/+ji7toby6glk2yzKvVg8jjOfjxr0PFbIyo5k43IYU+KfFOBR3i9pHbTbaLipbam4lAAtSC0ULThalhoGBUwtS21ICpYKEFqeKcCpAUbBtKRFNilSpI0WKWKVKiQYililSqEEBUwKVKoQralexabZvczK7IvcmM1xGsdIbvUsov6i2/u0PFh/EfsMUqVKzSaN+hxQk25Lo5LUz/aAoGAqiqlKlVodIXqG3kYjyqcE8ttLHPbyNHMh7DpzU01KrPsUje0bpNqtlcMWvZpourdurlO8Zxkc+P1rEuLu5uHkmluJXdzkszmlSo7Uizk+OQWMcKVKlQspZd0c/wDiMeeWGz8q6uw1nULIQj2lplYDCzdr5nn9aVKlyfI2MU0dZo+rJqZdOqMcqAFhnK/A1q4pqVMRnkqZMCpAU9KiVHxTgUqVQg+KlimpVCEwKlilSokP/9k="
                      }
                    />
                  </ListItemPrefix>
                  <div>
                    <Typography
                      variant='medium'
                      color='gray'
                      className={`font-normal bg-gray-300 p-2 -m-3 rounded-3xl ${
                        is_current_user
                          ? "bg-blue-500 text-white ml-auto mr-1"
                          : "bg-gray-300  mr-auto"
                      }`}
                    >
                      {content}
                    </Typography>
                  </div>
                </ListItem>
              );
            }
        )}
      </div>

      <div className='p-2 pb-5'>
        <div className='relative'>
          <input
            ref={inputRef}
            placeholder='Message...'
            onInput={handleTypeChange}
            value={message}
            onKeyDown={handleKeyPress}
            className='w-full focus:outline-none p-2 pr-10 rounded-3xl border-2 border-gray-300'
            type='text'
          />
          <button
            disabled={message === ""}
            onClick={handleSendMessage}
            className={`absolute right-2 top-2 ${
              message === "" ? "text-blue-300" : "text-blue-600"
            } px-4 rounded`}
          >
            Send
          </button>
        </div>
      </div>
    </div>
  );
};

export default SendMessages;
