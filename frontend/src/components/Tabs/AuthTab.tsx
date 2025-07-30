import { Box, Tabs, Tab } from "@mui/material";

interface AuthTabProps {
  value: number;
  onChange: (newValue: number) => void;
}

export default function AuthTab({ value, onChange }: AuthTabProps) {
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    onChange(newValue);
    console.log(event);
  };

  return (
    <Box
      sx={{
        borderBottom: 1,
        borderColor: "divider",
        width: "100%",
        backgroundColor: "black",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Tabs
        value={value}
        onChange={handleChange}
        aria-label="basic tabs example"
      >
        <Tab sx={{ color: "white" }} label="Sign In" />
        <Tab sx={{ color: "white" }} label="Sign Up" />
      </Tabs>
    </Box>
  );
}
