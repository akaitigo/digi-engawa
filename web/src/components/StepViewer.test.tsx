import { fireEvent, render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";
import type { Step } from "@/types/material";
import { StepViewer } from "./StepViewer";

const mockSteps: Step[] = [
	{
		id: "step-1",
		material_id: "mat-1",
		step_order: 1,
		title: "電源を入れる",
		body: "スマホの横のボタンを長押しします",
		furigana_body: "スマホのよこのボタンをながおしします",
		audio_text: "スマホの横のボタンを長押ししてください",
		created_at: "2026-01-01T00:00:00Z",
	},
	{
		id: "step-2",
		material_id: "mat-1",
		step_order: 2,
		title: "画面をタッチ",
		body: "画面の真ん中を指でタッチします",
		furigana_body: "画面(がめん)の真(ま)ん中(なか)を指(ゆび)でタッチします",
		audio_text: "画面の真ん中を指でタッチしてください",
		created_at: "2026-01-01T00:00:00Z",
	},
	{
		id: "step-3",
		material_id: "mat-1",
		step_order: 3,
		title: "パスコードを入れる",
		body: "数字を順番に押します",
		furigana_body: "数字(すうじ)を順番(じゅんばん)に押(お)します",
		audio_text: "数字を順番に押してください",
		created_at: "2026-01-01T00:00:00Z",
	},
];

describe("StepViewer", () => {
	it("shows the first step by default", () => {
		render(<StepViewer steps={mockSteps} />);
		expect(screen.getByText("電源を入れる")).toBeDefined();
		expect(screen.getByText("1 / 3 ステップ")).toBeDefined();
	});

	it("navigates to next step on button click", () => {
		render(<StepViewer steps={mockSteps} />);
		fireEvent.click(screen.getByText("つぎ →"));
		expect(screen.getByText("画面をタッチ")).toBeDefined();
		expect(screen.getByText("2 / 3 ステップ")).toBeDefined();
	});

	it("navigates to previous step", () => {
		render(<StepViewer steps={mockSteps} />);
		fireEvent.click(screen.getByText("つぎ →"));
		fireEvent.click(screen.getByText("← まえ"));
		expect(screen.getByText("電源を入れる")).toBeDefined();
	});

	it("disables prev button on first step", () => {
		render(<StepViewer steps={mockSteps} />);
		const prevButton = screen.getByText("← まえ");
		expect((prevButton as HTMLButtonElement).disabled).toBe(true);
	});

	it("disables next button on last step", () => {
		render(<StepViewer steps={mockSteps} />);
		fireEvent.click(screen.getByText("つぎ →"));
		fireEvent.click(screen.getByText("つぎ →"));
		const nextButton = screen.getByText("つぎ →");
		expect((nextButton as HTMLButtonElement).disabled).toBe(true);
	});

	it("calls onStepChange when navigating", () => {
		const onStepChange = vi.fn();
		render(<StepViewer steps={mockSteps} onStepChange={onStepChange} />);
		fireEvent.click(screen.getByText("つぎ →"));
		expect(onStepChange).toHaveBeenCalledWith(2);
	});

	it("shows empty message when no steps", () => {
		render(<StepViewer steps={[]} />);
		expect(screen.getByText("教材にステップがありません")).toBeDefined();
	});
});
