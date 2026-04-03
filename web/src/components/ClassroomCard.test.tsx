import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import type { Classroom } from "@/types/classroom";
import { ClassroomCard } from "./ClassroomCard";

const classroom: Classroom = {
	id: "c-1",
	title: "はじめてのスマホ教室",
	description: "スマートフォンの基本操作を学ぶ",
	location: "公民館A",
	capacity: 10,
	scheduled_at: "2026-05-01T10:00:00Z",
	classroom_code: "ABC123",
	created_at: "2026-01-01T00:00:00Z",
	updated_at: "2026-01-01T00:00:00Z",
};

describe("ClassroomCard", () => {
	it("displays classroom title", () => {
		render(<ClassroomCard classroom={classroom} />);
		expect(screen.getByText("はじめてのスマホ教室")).toBeDefined();
	});

	it("displays location", () => {
		render(<ClassroomCard classroom={classroom} />);
		expect(screen.getByText("公民館A")).toBeDefined();
	});

	it("displays classroom code", () => {
		render(<ClassroomCard classroom={classroom} />);
		expect(screen.getByText("コード: ABC123")).toBeDefined();
	});

	it("displays capacity", () => {
		render(<ClassroomCard classroom={classroom} />);
		expect(screen.getByText("定員 10名")).toBeDefined();
	});
});
