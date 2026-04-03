import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { ProgressBar } from "./ProgressBar";

describe("ProgressBar", () => {
	it("displays current/total steps", () => {
		render(<ProgressBar current={2} total={5} />);
		expect(screen.getByText("2 / 5 ステップ")).toBeDefined();
	});

	it("shows percentage", () => {
		render(<ProgressBar current={1} total={4} />);
		expect(screen.getByText("25%")).toBeDefined();
	});

	it("has proper aria attributes", () => {
		render(<ProgressBar current={3} total={10} />);
		const progressbar = screen.getByRole("progressbar");
		expect(progressbar.getAttribute("aria-valuenow")).toBe("30");
		expect(progressbar.getAttribute("aria-valuemin")).toBe("0");
		expect(progressbar.getAttribute("aria-valuemax")).toBe("100");
	});

	it("handles zero total", () => {
		render(<ProgressBar current={0} total={0} />);
		expect(screen.getByText("0%")).toBeDefined();
	});
});
