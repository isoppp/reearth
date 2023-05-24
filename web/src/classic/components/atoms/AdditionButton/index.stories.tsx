import { Story, Meta } from "@storybook/react";

import Component, { Props } from ".";

export default {
  title: "atoms/Buttons/AdditionButton",
  component: Component,
} as Meta;

export const Default: Story<Props> = args => <Component {...args} />;

Default.args = {};
