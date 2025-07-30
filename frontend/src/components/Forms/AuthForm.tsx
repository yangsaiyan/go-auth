import { Box, Button, TextField } from "@mui/material";
import { useState } from "react";
import { postData } from "../../utils/fetch";
import { useNavigate } from "react-router";

interface AuthFormProps {
  authTab: number;
}

export default function AuthForm({ authTab }: AuthFormProps) {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(formData);
    if (authTab === 0) {
      try {
        const response = await postData(
          "http://localhost:8080/api/auth/login",
          formData
        );
        console.log(response);
        navigate("/");
      } catch (error) {
        console.error(error);
      }
    } else {
      try {
        const response = await postData(
          "http://localhost:8080/api/auth/signup",
          formData
        );
        console.log(response);
      } catch (error) {
        console.error(error);
      }
    }
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      noValidate
      autoComplete="off"
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      gap={2}
      width={"fit-content"}
      height={"fit-content"}
      padding={2}
      borderRadius={2}
      border={1}
      borderColor="divider"
      sx={{ backgroundColor: "white" }}
    >
      <TextField
        sx={{ width: "100%", backgroundColor: "white" }}
        id="outlined-basic"
        label="Email"
        variant="outlined"
        name="email"
        value={formData.email}
        onChange={handleChange}
      />
      <TextField
        sx={{ width: "100%", backgroundColor: "white" }}
        id="outlined-basic"
        label="Password"
        variant="filled"
        name="password"
        value={formData.password}
        onChange={handleChange}
      />
      <Button
        type="submit"
        sx={{ width: "100%" }}
        variant="contained"
        color="primary"
      >
        {authTab === 0 ? "Sign In" : "Sign Up"}
      </Button>
    </Box>
  );
}
