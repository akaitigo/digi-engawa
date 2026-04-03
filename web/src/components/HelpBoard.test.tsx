import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import { HelpBoard } from "./HelpBoard";

describe("HelpBoard", () => {
	it("renders help board heading", () => {
		render(<HelpBoard classroomId="class-1" />);
		expect(screen.getByText("ヘルプボード")).toBeDefined();
	});

	it("shows empty state message", () => {
		render(<HelpBoard classroomId="class-1" />);
		expect(screen.getByText("まだヘルプリクエストはありません")).toBeDefined();
	});

	it("shows connection status", () => {
		render(<HelpBoard classroomId="class-1" />);
		expect(screen.getByText("○ 未接続")).toBeDefined();
	});
});
