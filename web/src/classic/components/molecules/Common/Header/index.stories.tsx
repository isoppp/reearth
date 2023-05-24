import { action } from "@storybook/addon-actions";
import { Meta, Story } from "@storybook/react";
import { BrowserRouter } from "react-router-dom";

import Component from "@reearth/classic/core/engines/Cesium";

import Header, { Props } from ".";

export default {
  title: "molecules/Common/Header",
  component: Header,
  args: {
    user: {
      name: "Shinnosuke Komiya",
    },
    currentWorkspace: {
      id: "1",
      name: "Darwin Education",
    },
    workspaces: [
      {
        id: "A",
        name: "Team A",
      },
      {
        id: "B",
        name: "Team B",
      },
    ],
    onSignOut: () => action("signOut"),
  },
} as Meta<Props>;

const Template: Story<Props> = args => <BrowserRouter><Component {...args} /></BrowserRouter>;

export const Default = Template.bind({});

export const NoLogin = Template.bind({
  onSignOut: undefined,
  user: undefined,
});
