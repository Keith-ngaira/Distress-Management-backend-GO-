import React from 'react';
import { useNavigate } from 'react-router-dom';
import {
  AppBar,
  Toolbar,
  Typography,
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Box,
  Container,
  IconButton,
} from '@mui/material';
import {
  Dashboard as DashboardIcon,
  Add as AddIcon,
  Brightness4 as DarkIcon,
  Brightness7 as LightIcon,
} from '@mui/icons-material';
import { useThemeMode } from '../contexts/ThemeContext';

const drawerWidth = 240;

function Layout({ children }) {
  const navigate = useNavigate();
  const { mode, toggleColorMode } = useThemeMode();

  const menuItems = [
    { text: 'Cases Dashboard', icon: <DashboardIcon />, path: '/' },
    { text: 'New Distress Case', icon: <AddIcon />, path: '/new-case' },
  ];

  return (
    <Box sx={{ display: 'flex' }}>
      <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between' }}>
          <Typography variant="h6" noWrap component="div">
            Distress Messages Management System
          </Typography>
          <IconButton
            color="inherit"
            onClick={toggleColorMode}
            aria-label="toggle theme"
          >
            {mode === 'dark' ? <LightIcon /> : <DarkIcon />}
          </IconButton>
        </Toolbar>
      </AppBar>
      <Drawer
        variant="permanent"
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
      >
        <Toolbar />
        <Box sx={{ overflow: 'auto' }}>
          <List>
            {menuItems.map((item) => (
              <ListItem
                button
                key={item.text}
                onClick={() => navigate(item.path)}
              >
                <ListItemIcon>{item.icon}</ListItemIcon>
                <ListItemText primary={item.text} />
              </ListItem>
            ))}
          </List>
        </Box>
      </Drawer>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          p: 3,
          width: { sm: `calc(100% - ${drawerWidth}px)` },
        }}
      >
        <Toolbar />
        <Container>{children}</Container>
      </Box>
    </Box>
  );
}

export default Layout;
