"use client";

import React, { ChangeEvent, useState } from "react";
import TitleHeader from "./title-header";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";

export const ContactSection = () => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    message: "",
  });

  const { name, email, message } = formData;

  const isDiasbled = !name || !message || !email;

  const handleChange = (
    e: ChangeEvent<HTMLInputElement & HTMLTextAreaElement>
  ) => {
    const { id, value } = e.target;
    setFormData((prev) => {
      return {
        ...prev,
        [id]: value,
      };
    });
  };

  return (
    <section className="p-6 space-y-5">
      <TitleHeader
        title="Contact Us"
        subtitle="We'd love to hear from you. drop a message and we'll get back soon."
      />

      <div className="max-w-2xl mx-auto">
        <form
          action={"https://app.proforms.top/f/pr84642fa"}
          className="space-y-4"
        >
          <div className="space-y-2.5">
            <Label htmlFor="name">Name</Label>
            <Input
              id="name"
              value={name}
              onChange={handleChange}
              placeholder="Your name"
            />
          </div>

          <div className="space-y-2.5">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              value={email}
              onChange={handleChange}
              placeholder="you@example.com"
            />
          </div>

          <div className="space-y-2.5">
            <Label htmlFor="message">Message</Label>
            <Textarea
              id="message"
              value={message}
              className="max-h-[400px]"
              onChange={handleChange}
              placeholder="Write your message..."
              rows={6}
            />
          </div>

          <div className="pt-2">
            <Button type="submit" disabled={isDiasbled}>
              Send Message
            </Button>
          </div>
        </form>
      </div>
    </section>
  );
};

export default ContactSection;
