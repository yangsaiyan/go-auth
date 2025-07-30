import { Button, Grid } from "@mui/material";
import { useNavigate } from "react-router";
import { postData } from "../../utils/fetch";
import { useEffect } from "react";
import { checkAuth } from "../../utils/check-auth";

export default function Home() {
  const navigate = useNavigate();
  const handleLogout = async () => {
    try {
      const response = await postData(
        "http://localhost:8080/api/auth/logout",
        {}
      );
      console.log(response);
      navigate("/auth");
    } catch (error) {
      console.log(error);
      console.error(error);
    }
  };

  useEffect(() => {
    checkAuth().then((res) => {
      if (!res) {
        navigate("/auth");
      }
    });
  }, []);

  return (
    <Grid
      container
      width="100vw"
      height="100vh"
      justifyContent="center"
      alignItems="center"
    >
      <Button
        sx={{ backgroundColor: "red", color: "white" }}
        onClick={handleLogout}
      >
        Logout
      </Button>
    </Grid>
  );
}
