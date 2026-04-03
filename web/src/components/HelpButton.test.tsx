import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { HelpButton } from "./HelpButton";

describe("HelpButton", () => {
	it("renders with help text", () => {
		render(<HelpButton classroomId="c1" participantId="p1" materialStepId="s1" />);
		expect(screen.getByText("🆘 こまった！")).toBeDefined();
	});

	it("has accessible label", () => {
		render(<HelpButton classroomId="c1" participantId="p1" materialStepId="s1" />);
		expect(screen.getByRole("button", { name: "助けを呼ぶ" })).toBeDefined();
	});

	it("has minimum touch target", () => {
		render(<HelpButton classroomId="c1" participantId="p1" materialStepId="s1" />);
		const button = screen.getByRole("button");
		expect(button.style.minHeight).toBe("80px");
	});
});
