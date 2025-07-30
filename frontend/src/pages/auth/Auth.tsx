import { Grid } from "@mui/material";
import AuthTab from "../../components/Tabs/AuthTab";
import { useState } from "react";
import AuthForm from "../../components/Forms/AuthForm";

export default function Auth() {
  const [value, setValue] = useState(0);

  return (
    <Grid
      container
      justifyContent="center"
      alignItems="center"
      flexDirection="column"
      width="100vw"
      height="100vh"
      sx={{ backgroundColor: "black" }}
    >
      <AuthTab value={value} onChange={setValue} />
      <AuthForm authTab={value} />
    </Grid>
  );
}
