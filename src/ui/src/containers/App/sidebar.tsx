import * as React from 'react';
import Drawer from '@material-ui/core/Drawer';
import {
  createStyles, WithStyles, withStyles, Theme,
} from '@material-ui/core/styles';

import ClusterIcon from 'components/icons/cluster';
import NamespaceIcon from 'components/icons/namespace';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import SettingsIcon from 'components/icons/settings';
import { Link } from 'react-router-dom';
import Tooltip from '@material-ui/core/Tooltip';
import ClusterContext from 'common/cluster-context';
import { toEntityPathname, LiveViewPage } from 'components/live-widgets/utils/live-view-params';
import { Divider } from '@material-ui/core';

const styles = (
  {
    spacing,
    palette,
    transitions,
    breakpoints,
  }: Theme) => createStyles({
  drawerOpen: {
    width: spacing(29),
    zIndex: 1250,
    flexShrink: 0,
    whiteSpace: 'nowrap',
    paddingTop: spacing(8),
    transition: transitions.create('width', {
      easing: transitions.easing.sharp,
      duration: transitions.duration.enteringScreen,
    }),
    overflowX: 'hidden',
    backgroundColor: palette.foreground.grey3,
    boxShadow: `${spacing(0.25)}px 0px ${spacing(2)}px ${palette.background.five}`,
    paddingBottom: spacing(2),
  },
  drawerClose: {
    transition: transitions.create('width', {
      easing: transitions.easing.sharp,
      duration: transitions.duration.leavingScreen,
    }),
    width: spacing(6),
    zIndex: 1250,
    overflowX: 'hidden',
    paddingTop: spacing(8),
    backgroundColor: palette.foreground.grey3,
    boxShadow: `${spacing(0.25)}px 0px ${spacing(2)}px ${palette.background.five}`,
    paddingBottom: spacing(2),
    [breakpoints.down('sm')]: {
      display: 'none',
    },
  },
  docked: {
    position: 'absolute',
  },
  spacer: {
    flex: 1,
  },
  listIcon: {
    paddingLeft: spacing(1.5),
    paddingTop: spacing(1),
    paddingBottom: spacing(1),
  },
  divider: {
    backgroundColor: palette.foreground.grey2,
  },
  profile: {
    padding: 0,
    marginLeft: -1 * spacing(1.5),
  },
  profileText: {
    whiteSpace: 'nowrap',
    '& .MuiTypography-displayBlock': {
      overflow: 'hidden',
      textOverflow: 'ellipsis',
    },
  },
});

interface SideBarProps extends WithStyles<typeof styles> {
  open: boolean;
}

const profileItems = [
  {
    icon: <SettingsIcon />,
    link: '/admin',
    text: 'Admin',
  },
];

const SideBarItem = ({
  classes, icon, link, text,
}) => (
  <Tooltip title={text}>
    <ListItem button component={Link} to={link} key={text} className={classes.listIcon}>
      <ListItemIcon>{icon}</ListItemIcon>
      <ListItemText primary={text} />
    </ListItem>
  </Tooltip>
);

const SideBar = ({
  classes, open,
}) => {
  const { selectedClusterName } = React.useContext(ClusterContext);

  const navItems = React.useMemo(() => (
    [{
      icon: <ClusterIcon />,
      link: toEntityPathname({ params: {}, clusterName: selectedClusterName, page: LiveViewPage.Cluster }),
      text: 'Cluster',
    },
    {
      icon: <NamespaceIcon />,
      link: toEntityPathname({ params: {}, clusterName: selectedClusterName, page: LiveViewPage.Namespaces }),
      text: 'Namespaces',
    }]
  ), [selectedClusterName]);

  return (
    <Drawer
      variant='permanent'
      className={open ? classes.drawerOpen : classes.drawerClose}
      classes={{
        paper: open ? classes.drawerOpen : classes.drawerClose,
        docked: classes.docked,
      }}
    >
      <List>
        {navItems.map(({ icon, link, text }) => (
          <SideBarItem key={text} classes={classes} icon={icon} link={link} text={text} />
        ))}
      </List>
      <div className={classes.spacer} />
      <Divider />
      <List>
        {profileItems.map(({ icon, link, text }) => (
          <SideBarItem key={text} classes={classes} icon={icon} link={link} text={text} />
        ))}
      </List>
    </Drawer>
  );
};

export default withStyles(styles)(SideBar);