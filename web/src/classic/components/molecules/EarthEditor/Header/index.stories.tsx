import { action } from "@storybook/addon-actions";
import {Meta, Story} from "@storybook/react";
import { BrowserRouter } from "react-router-dom";

import Header from ".";
import Component, {Props} from "@reearth/classic/core/Visualizer";

const defaultProps = {
  currentProject: {
    id: "1",
    name: "Sample Project",
  },
  currentProjectStatus: "published" as const,
  recentProjects: [
    {
      id: "1",
      name: "Project 1",
    },
    {
      id: "2",
      name: "Project 2",
    },
    {
      id: "3",
      name: "Project 3",
    },
    {
      id: "4",
      name: "Project 4",
    },
    {
      id: "5",
      name: "Project 5",
    },
  ],
  user: {
    name: "Shinnosuke Komiya",
  },
  currentTeam: {
    id: "1",
    name: "Darwin Education",
  },
  teams: [
    {
      id: "A",
      name: "Team A",
    },
    {
      id: "B",
      name: "Team B",
    },
  ],
  onBack: () => action("onBack"),
  onForward: () => action("onForward"),
  onSignOut: () => action("signOut"),
};

export default {
  title: "molecules/EarthEditor/Header",
  component: Header,
} as Meta;

const Template: Story<Props> = args => (
  <BrowserRouter>
    <Component {...args} />
  </BrowserRouter>
);

export const Default = () => Template.bind({});
