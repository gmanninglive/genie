import { render } from "@testing-library/react";

import { {{ toTitle name }} } from "./{{ toTitle name }}";


describe("{{ toTitle name }} Component", () => {
  it("renders correctly without throwing exception", () => {
    const { container } = render(
      <{{ toTitle name }} />
    );

    expect(container).toBeInTheDocument()
  });