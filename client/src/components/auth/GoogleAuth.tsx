import { GoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import UserAuth from "../../api/auth/user";

function GoogleAuthComponent(): JSX.Element {
  const navigate = useNavigate();
  const userAuth = UserAuth();

  const errorMessage = (): void => {
    toast.error("Failed to login by google,Please try again later", {
      position: toast.POSITION.BOTTOM_RIGHT,
    });
  };

  const handleSignInWithGoogle = async (credential: string) => {
    userAuth
      .googleSingIn(credential)
      .then((response) => {
        toast.success(response?.message, {
          position: toast.POSITION.BOTTOM_RIGHT,
        });
        response?.success&&navigate('/')
      })
      .catch((error) => {
        toast.error(error?.data?.error[0], {
          position: toast.POSITION.BOTTOM_RIGHT,
        });
      });
  };

  return (
    <div className='mb-5 '>
      <div className='flex justify-center'>
        <GoogleLogin
          width='352px'
          size='large'
          // theme="filled_blue"
          logo_alignment='center'
          shape='pill'
          auto_select={false}
          type='standard'
          ux_mode='popup'
          onSuccess={(response) => {
            if (response) {
              handleSignInWithGoogle(response.credential ?? "empty response");
            }
          }}
          onError={errorMessage}
        />
      </div>
    </div>
  );
}

export default GoogleAuthComponent;
