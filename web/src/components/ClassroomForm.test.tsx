import { render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import { ClassroomForm } from "./ClassroomForm";

describe("ClassroomForm", () => {
	it("renders form fields", () => {
		render(<ClassroomForm onSubmit={vi.fn()} />);
		expect(screen.getByLabelText("教室名 *")).toBeDefined();
		expect(screen.getByLabelText("説明")).toBeDefined();
		expect(screen.getByLabelText("場所")).toBeDefined();
		expect(screen.getByLabelText("日付 *")).toBeDefined();
		expect(screen.getByLabelText("時間")).toBeDefined();
		expect(screen.getByLabelText("定員")).toBeDefined();
	});

	it("renders submit button", () => {
		render(<ClassroomForm onSubmit={vi.fn()} />);
		expect(screen.getByText("教室を作成する")).toBeDefined();
	});

	it("shows submitting state", () => {
		render(<ClassroomForm onSubmit={vi.fn()} submitting={true} />);
		expect(screen.getByText("作成中...")).toBeDefined();
	});
});
